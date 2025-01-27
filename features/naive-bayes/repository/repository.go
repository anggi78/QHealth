package repository

import (
	"fmt"
	"qhealth/domain"
	naivebayes "qhealth/features/naive-bayes"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewNaiveRepository(db *gorm.DB) naivebayes.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SavePatientsToDB(patients []domain.Patient) error {
    for _, patient := range patients {
        var existingPatient domain.Patient
        result := r.db.Where("name = ? AND age = ?", patient.Name, patient.Age).First(&existingPatient)

        if result.Error == nil { 

            existingPatient.Diagnosis = patient.Diagnosis
            existingPatient.Category = patient.Category
            existingPatient.Priority = patient.Priority

            if err := r.db.Save(&existingPatient).Error; err != nil {
                return fmt.Errorf("failed to update patient %s: %v", patient.Name, err)
            }
        } else { 
            if err := r.db.Create(&patient).Error; err != nil {
                return fmt.Errorf("failed to insert patient %s: %v", patient.Name, err)
            }
        }
    }

    return nil
}

func (r *repository) GetAllPatients() ([]domain.Patient, error) {
	var patients []domain.Patient
	err := r.db.Find(&patients).Error
    if err != nil {
        return nil, err
    }
	return patients, nil
}
