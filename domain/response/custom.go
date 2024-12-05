package response

import "time"

type GetStockUpdationResponse struct {
	ID          int                        `json:"id" gorm:"column:id"`
	IsAddtion   bool                       `json:"is_addition" gorm:"column:is_addition"`
	BroughtAt   string                     `json:"brought_at" gorm:"column:brought_at"`
	Particulars []StockUpdationParticulars `json:"particulars" gorm:"-"`
}

type MedicineWiseStockUpdationDetails struct {
	StockUpdationID int       `json:"stock_updation_id" gorm:"column:stock_updation_id"`
	BroughtAt       time.Time `json:"brought_at" gorm:"column:brought_at"`
	Quantity        int       `json:"quantity" gorm:"column:quantity"`
}

type StockUpdationParticulars struct {
	MedicineID int `json:"medicine_id" gorm:"column:medicine_id"`
	Quantity   int `json:"quantity" gorm:"column:quantity"`
}
