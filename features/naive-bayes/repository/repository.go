package repository

import (
	"log"
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
            log.Printf("Patient %s (Age: %d) already exists in the database, skipping...", patient.Name, patient.Age)
            continue
        }

        if err := r.db.Create(&patient).Error; err != nil {
            log.Printf("Failed to insert patient: %+v, error: %v", patient, err)
            return err
        }
    }

    log.Println("All patients processed successfully")
    return nil
}



func (r *repository) GetAllPatients() ([]domain.Patient, error) {
	var patients []domain.Patient
	if err := r.db.Raw("SELECT id, name, age, diagnosis, category FROM patients").Scan(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
