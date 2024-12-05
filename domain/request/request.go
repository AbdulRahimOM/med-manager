package request

import (
	"med-manager/models"
	"time"

	"gorm.io/gorm"
)

type MedicineRequest struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	TypeID       int     `json:"typeId" validate:"gte=1"`
	Price        float64 `json:"price" validate:"gte=0"`
	MinStock     int     `json:"min_stock" validate:"gte=0"`
	OptimalStock int     `json:"optimal_stock" validate:"gte=0"`
}

func (m *MedicineRequest) ToMedicine() *models.Medicine {
	return &models.Medicine{
		Name:         m.Name,
		Description:  m.Description,
		TypeID:       m.TypeID,
		Price:        m.Price,
		MinStock:     m.MinStock,
		OptimalStock: m.OptimalStock,
	}
}

type PatientReq struct {
	Name        string         `json:"name" gorm:"column:name" validate:"required"`
	Age         int            `json:"age" gorm:"column:age"`
	Gender      string         `json:"gender" gorm:"column:gender"`
	Contact     string         `json:"contact" gorm:"column:contact"`
	Description string         `json:"description" gorm:"column:description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *PatientReq) ToPatient() *models.Patient {
	return &models.Patient{
		Name:        p.Name,
		Age:         p.Age,
		Gender:      p.Gender,
		Contact:     p.Contact,
		Description: p.Description,
	}
}

type VisitReq struct {
	PatientID int       `json:"patient_id" gorm:"column:patient_id" validate:"required,gte=1"`
	Date      time.Time `json:"date" gorm:"column:date"`
	Notes     string    `json:"notes" gorm:"column:notes"`
}

func (v *VisitReq) ToVisit() *models.Visit {
	return &models.Visit{
		PatientID: v.PatientID,
		Date:      v.Date,
		Notes:     v.Notes,
	}
}
