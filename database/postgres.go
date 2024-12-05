package database

import (
	"log"
	models "med-manager/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = InitDB()
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=medical_store port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	err = db.AutoMigrate(
		&models.Medicine{},
		&models.MedType{},
		&models.StockUpdation{},
		&models.StockUpdationParticulars{},
		&models.Patient{},
		&models.Visit{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
