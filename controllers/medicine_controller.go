package controllers

import (
	"gorm.io/gorm"
)

type MedicineController struct {
	DB *gorm.DB
}

func NewMedicineController(db *gorm.DB) *MedicineController {
	return &MedicineController{DB: db}
}
