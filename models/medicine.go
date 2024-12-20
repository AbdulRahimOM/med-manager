package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

var ErrUniqueNameViolation = fmt.Errorf("Name already exists")

type Medicine struct {
	ID           int       `json:"id" gorm:"column:id;primaryKey"`
	Name         string    `json:"name" gorm:"column:name;unique" validate:"required"`
	Description  string    `json:"description" gorm:"column:description"`
	TypeID       int       `json:"typeId" gorm:"column:type_id" validate:"required,gte=1"`
	Price        float64   `json:"price" gorm:"column:price" validate:"required,gte=0"`
	MinStock     int       `json:"min_stock" gorm:"column:min_stock" validate:"required,gte=0"`
	OptimalStock int       `json:"optimal_stock" gorm:"column:optimal_stock" validate:"required,gte=0"`
	CurrentStock int       `json:"current_stock" gorm:"column:current_stock;default:0" validate:"gte=0"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`

	Type MedType `json:"-" gorm:"foreignKey:TypeID;references:ID"`
}

type MedType struct {
	ID   int    `json:"id" gorm:"column:id;primaryKey"`
	Type string `json:"type" gorm:"column:type;unique" validate:"required"`
}

// Model methods for database operations
func (m *Medicine) Create(db *gorm.DB) error {
	err := db.Create(m).Error
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_medicines_name\" (SQLSTATE 23505)" {
			return ErrUniqueNameViolation
		} else {
			return err
		}
	}
	return nil
}

func (m *Medicine) Update(db *gorm.DB) error {
	err := db.Save(m).Error
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_medicines_name\" (SQLSTATE 23505)" {
			return ErrUniqueNameViolation
		} else {
			return err
		}
	}
	return nil
}

func GetMedicineByID(db *gorm.DB, id int) (*Medicine, error) {
	var medicine Medicine
	err := db.First(&medicine, id).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func GetAllMedicines(db *gorm.DB) ([]Medicine, error) {
	var medicines []Medicine
	err := db.Find(&medicines).Error
	return medicines, err
}

func DeleteMedicine(db *gorm.DB, id int) error {
	return db.Delete(&Medicine{}, id).Error
}

func GetAllMedTypes(db *gorm.DB) ([]MedType, error) {
	var medTypes []MedType
	err := db.Find(&medTypes).Error
	return medTypes, err
}

func GetMedTypeByID(db *gorm.DB, id int) (*MedType, error) {
	var medType MedType
	err := db.First(&medType, id).Error
	if err != nil {
		return nil, err
	}
	return &medType, nil
}

func (m *MedType) Create(db *gorm.DB) error {
	m.ID = 0 //To prevent id from being set by the client
	err := db.Create(m).Error
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_med_types_type\" (SQLSTATE 23505)" {
			return ErrUniqueNameViolation
		} else {
			return err
		}
	}
	return nil
}

func (m *MedType) Update(db *gorm.DB) error {
	result := db.Exec("UPDATE med_types SET type = ? WHERE id = ?", m.Type, m.ID)
	if result.Error != nil {
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"uni_med_types_type\" (SQLSTATE 23505)" {
			return ErrUniqueNameViolation
		}
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("there is no such id")
	}
	return nil
}

func DeleteMedType(db *gorm.DB, id int) error {
	return db.Delete(&MedType{}, id).Error
}
