package domain

import (
	"qhealth/helpers"

	"gorm.io/gorm"
)

type Patient struct {
	Id        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Diagnosis string `json:"diagnosis"`
	Category  string `json:"category"`
	Priority  string `json:"priority"`
}

type PatientResp struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Diagnosis string `json:"diagnosis"`
	Category  string `json:"category"`
	Priority  string `json:"priority"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	p.Id = helpers.CreateId()
	return nil
}