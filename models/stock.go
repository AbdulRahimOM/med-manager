package models

import (
	"fmt"
	"med-manager/domain/response"
	"time"

	"gorm.io/gorm"
)

var ErrInsufficientStock = fmt.Errorf("Insufficient stock")

func (sReq *StockUpdateRequest) AddToStock(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	stockUpdation := &StockUpdation{
		BroughtAt: time.Now(),
		IsAddtion: true,
	}
	err := tx.Create(stockUpdation).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, stockChange := range sReq.StockChanges {
		stockUpdationParticulars := &StockUpdationParticulars{
			StockUpdationID: stockUpdation.ID,
			MedicineID:      stockChange.MedicineID,
			Quantity:        stockChange.Quantity,
		}
		err := tx.Create(stockUpdationParticulars).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		//add stockChange.Quantity to Medicine.CurrentStock
		var medicine Medicine
		err = tx.Model(&medicine).Where("id = ?", stockChange.MedicineID).Update("current_stock", gorm.Expr("current_stock + ?", stockChange.Quantity)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (sReq *StockUpdateRequest) DeductFromStock(db *gorm.DB) (error, int) {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error, 0
	}

	stockUpdation := &StockUpdation{
		BroughtAt: time.Now(),
		IsAddtion: false,
	}
	err := tx.Create(stockUpdation).Error
	if err != nil {
		tx.Rollback()
		return err, 0
	}

	for _, stockChange := range sReq.StockChanges {
		stockUpdationParticulars := &StockUpdationParticulars{
			StockUpdationID: stockUpdation.ID,
			MedicineID:      stockChange.MedicineID,
			Quantity:        stockChange.Quantity,
		}
		err := tx.Create(stockUpdationParticulars).Error
		if err != nil {
			tx.Rollback()
			return err, 0
		}

		//get current stock
		var currentStock int
		err = tx.Raw("SELECT current_stock FROM medicines WHERE id = ?", stockChange.MedicineID).Scan(&currentStock).Error
		if err != nil {
			tx.Rollback()
			return err, 0
		}

		if currentStock < stockChange.Quantity {
			tx.Rollback()
			return ErrInsufficientStock, stockChange.MedicineID
		}

		//deduct stockChange.Quantity from Medicine.CurrentStock
		var medicine Medicine
		err = tx.Model(&medicine).Where("id = ?", stockChange.MedicineID).Update("current_stock", gorm.Expr("current_stock - ?", stockChange.Quantity)).Error
		if err != nil {
			tx.Rollback()
			return err, 0
		}
	}

	return tx.Commit().Error, 0
}

func GetAllStockUpdations(db *gorm.DB, isAddtion bool, offset, limit int) ([]response.GetStockUpdationResponse, error) {
	var stockAdditions []response.GetStockUpdationResponse
	err := db.Table("stock_updations").Where("is_addition = ?", isAddtion).Offset(offset).Limit(limit).Find(&stockAdditions).Error
	if err != nil {
		return nil, err
	}

	for i := range stockAdditions {
		err := db.Table("stock_updation_particulars").Where("stock_updation_id = ?", stockAdditions[i].ID).Find(&stockAdditions[i].Particulars).Error
		if err != nil {
			return nil, err
		}
	}

	return stockAdditions, nil
}

func GetStockUpdationByID(db *gorm.DB, id int) (*response.GetStockUpdationResponse, error) {
	var stockAddition response.GetStockUpdationResponse
	err := db.Table("stock_updations").Where("id = ?", id).First(&stockAddition).Error
	if err != nil {
		return nil, err
	}

	err = db.Table("stock_updation_particulars").Where("stock_updation_id = ?", stockAddition.ID).Find(&stockAddition.Particulars).Error
	if err != nil {
		return nil, err
	}

	return &stockAddition, nil
}

func DeleteStockUpdation(db *gorm.DB, id int) error {
	var stockUpdation StockUpdation
	err := db.Where("id = ?", id).First(&stockUpdation).Error
	if err != nil {
		return err
	}

	var stockUpdationParticulars []StockUpdationParticulars
	err = db.Where("stock_updation_id = ?", id).Find(&stockUpdationParticulars).Error
	if err != nil {
		db.Rollback()
		return err
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	fmt.Println("stockUpdationParticulars: ", stockUpdationParticulars)
	for i := range stockUpdationParticulars {
		// adjusting stock to undo the stock updation
		var stockChangeToDo int
		if stockUpdation.IsAddtion {
			stockChangeToDo = -stockUpdationParticulars[i].Quantity
		} else {
			stockChangeToDo = stockUpdationParticulars[i].Quantity
		}
		var medicine Medicine
		err = tx.Model(&medicine).Where("id = ?", stockUpdationParticulars[i].MedicineID).Update("current_stock", gorm.Expr("current_stock + ?", stockChangeToDo)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Delete(&StockUpdation{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteStockUpdationParticulars(db *gorm.DB, stockUpdationID, medicineID int) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	//deduct stockChange.Quantity from Medicine.CurrentStock
	var stockUpdationParticular StockUpdationParticulars
	err := tx.Where("stock_updation_id = ? AND medicine_id = ?", stockUpdationID, medicineID).First(&stockUpdationParticular).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var medicine Medicine
	err = tx.Model(&medicine).Where("id = ?", stockUpdationParticular.MedicineID).Update("current_stock", gorm.Expr("current_stock - ?", stockUpdationParticular.Quantity)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&StockUpdationParticulars{}, "stock_updation_id = ? AND medicine_id = ?", stockUpdationID, medicineID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetStockUpdationParticularsByStockUpdationID(db *gorm.DB, stockUpdationID int) ([]StockUpdationParticulars, error) {
	var stockUpdationParticulars []StockUpdationParticulars
	err := db.Where("stock_updation_id = ?", stockUpdationID).Find(&stockUpdationParticulars).Error
	if err != nil {
		return nil, err
	}

	return stockUpdationParticulars, nil
}

func GetStockUpdationParticularsByMedicineID(db *gorm.DB, medicineID int, isAddition bool) ([]response.MedicineWiseStockUpdationDetails, error) {
	var stockUpdationParticulars []response.MedicineWiseStockUpdationDetails
	query := `
		SELECT
			sup.stock_updation_id,
			su.brought_at,
			sup.quantity
		FROM
			stock_updation_particulars sup
		JOIN
			stock_updations su
		ON
			sup.stock_updation_id = su.id
		WHERE
			sup.medicine_id = ?
			AND su.is_addition = ?
	`
	err := db.Raw(query, medicineID, isAddition).Scan(&stockUpdationParticulars).Error
	if err != nil {
		return nil, err
	}

	return stockUpdationParticulars, nil
}

func UpdateParticularsInAnStockUpdation(db *gorm.DB, stockUpdationID int, stockChanges []StockChanges) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var oldParticulars []StockUpdationParticulars
	err := tx.Where("stock_updation_id = ?", stockUpdationID).Find(&oldParticulars).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	oldParticularsMap := make(map[int]int)
	for i := range oldParticulars {
		oldParticularsMap[oldParticulars[i].MedicineID] = oldParticulars[i].Quantity
	}

	for i := range stockChanges {
		quantity, ok := oldParticularsMap[stockChanges[i].MedicineID]
		if !ok {
			//add new particular
			stockUpdationParticulars := &StockUpdationParticulars{
				StockUpdationID: stockUpdationID,
				MedicineID:      stockChanges[i].MedicineID,
				Quantity:        stockChanges[i].Quantity,
			}
			err := tx.Create(stockUpdationParticulars).Error
			if err != nil {
				tx.Rollback()
				return err
			}

			//add stockChange.Quantity to Medicine.CurrentStock
			var medicine Medicine
			err = tx.Model(&medicine).Where("id = ?", stockChanges[i].MedicineID).Update("current_stock", gorm.Expr("current_stock + ?", stockChanges[i].Quantity)).Error
			if err != nil {
				tx.Rollback()
				return err
			}

		} else {
			//update particular
			err := tx.Model(&StockUpdationParticulars{}).Where("stock_updation_id = ? AND medicine_id = ?", stockUpdationID, stockChanges[i].MedicineID).Update("quantity", stockChanges[i].Quantity).Error
			if err != nil {
				tx.Rollback()
				return err
			}

			//update Medicine.CurrentStock
			var medicine Medicine
			err = tx.Model(&medicine).Where("id = ?", stockChanges[i].MedicineID).Update("current_stock", gorm.Expr("current_stock + ?", stockChanges[i].Quantity-quantity)).Error
			if err != nil {
				tx.Rollback()
				return err
			}

			delete(oldParticularsMap, stockChanges[i].MedicineID)
		}
	}

	for medicineID, quantity := range oldParticularsMap {
		//delete particular
		err := tx.Delete(&StockUpdationParticulars{}, "stock_updation_id = ? AND medicine_id = ?", stockUpdationID, medicineID).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		//deduct stockChange.Quantity from Medicine.CurrentStock
		var medicine Medicine
		err = tx.Model(&medicine).Where("id = ?", medicineID).Update("current_stock", gorm.Expr("current_stock - ?", quantity)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func GetMedicineStockByMedicineID(db *gorm.DB, medicineID int) (int, error) {
	var currentStock int
	err := db.Raw("SELECT current_stock FROM medicines WHERE id = ?", medicineID).Scan(&currentStock).Error
	if err != nil {
		return 0, err
	}

	return currentStock, nil
}
