package models

import "time"

type StockUpdation struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
	IsAddtion bool      `json:"is_addition" gorm:"column:is_addition"`
	BroughtAt time.Time `json:"brought_at" gorm:"column:brought_at"`
}

func (s *StockUpdation) TableName() string {
	return "stock_updations"
}

type StockUpdationParticulars struct {
	StockUpdationID int `json:"stock_updation_id" gorm:"column:stock_updation_id;primaryKey"`
	MedicineID      int `json:"medicine_id" gorm:"column:medicine_id;primaryKey"`
	Quantity        int `json:"quantity" gorm:"column:quantity;primaryKey"`

	Medicine      Medicine      `json:"-" gorm:"foreignKey:MedicineID;references:ID"`
	StockUpdation StockUpdation `json:"-" gorm:"foreignKey:StockUpdationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *StockUpdationParticulars) TableName() string {
	return "stock_updation_particulars"
}
