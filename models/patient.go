package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID          int            `json:"id" gorm:"column:id;primaryKey"`
	Name        string         `json:"name" gorm:"column:name" validate:"required"`
	Age         int            `json:"age" gorm:"column:age"`
	Gender      string         `json:"gender" gorm:"column:gender"`
	Contact     string         `json:"contact" gorm:"column:contact"`
	Description string         `json:"description" gorm:"column:description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Patient) Create(db *gorm.DB) error {
	p.ID = 0 //to prevent id from being set by the client
	p.CreatedAt = time.Now()
	return db.Create(p).Error
}

func (p *Patient) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func GetPatientByID(db *gorm.DB, id int) (*Patient, error) {
	var patient Patient
	err := db.First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func GetAllPatients(db *gorm.DB, offset, limit int) ([]Patient, error) {
	var patients []Patient
	err := db.Raw("SELECT * FROM patients WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset).Scan(&patients).Error
	return patients, err
}

func DeletePatient(db *gorm.DB, id int) error {
	return db.Delete(&Patient{}, id).Error
}

type Visit struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
	PatientID int       `json:"patient_id" gorm:"column:patient_id"`
	Date      time.Time `json:"date" gorm:"column:date"`
	Notes     string    `json:"notes" gorm:"column:notes"`

	Patient Patient `json:"-" gorm:"foreignKey:PatientID;references:ID"`
}

func (v *Visit) Create(db *gorm.DB) error {
	return db.Create(v).Error
}

func (v *Visit) Update(db *gorm.DB) error {
	return db.Save(v).Error
}

func GetVisitByID(db *gorm.DB, id int) (*Visit, error) {
	var visit Visit
	err := db.First(&visit, id).Error
	if err != nil {
		return nil, err
	}
	return &visit, nil
}

func GetAllVisits(db *gorm.DB,offset,limit int) ([]Visit, error) {
	var visits []Visit
	err := db.Raw("SELECT * FROM visits ORDER BY date DESC LIMIT ? OFFSET ?", limit, offset).Scan(&visits).Error
	return visits, err
}

func DeleteVisit(db *gorm.DB, id int) error {
	return db.Delete(&Visit{}, id).Error
}

func GetAllVisitsByPatientID(db *gorm.DB, patientID int) ([]Visit, error) {
	var visits []Visit
	err := db.Where("patient_id = ?", patientID).Find(&visits).Error
	return visits, err
}
