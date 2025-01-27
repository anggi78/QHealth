package service

import (
	"log"
	"qhealth/domain"
	naivebayes "qhealth/features/naive-bayes"
	"qhealth/features/naive-bayes/nb"
	"qhealth/helpers/excel"
	"strings"
)

type service struct {
	repo naivebayes.Repository
}

func NewNaiveService(repo naivebayes.Repository) naivebayes.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllPatients() ([]domain.PatientResp, error) {
	patients, err := s.repo.GetAllPatients()
	if err != nil {
		return nil, err
	}

	result := domain.ListPatientToResp(patients)
	return result, nil
}

func (s *service) GetPatientsByPriority(priority string) ([]domain.PatientResp, error) {
    patients, err := s.repo.GetPatientsByPriority(priority)
    if err != nil {
        return nil, err
    }

    result := domain.ListPatientToResp(patients)
    return result, nil
}


func (s *service) ImportPatientsFromExcel(filePath string) error {
	patients, err := excel.ReadExcelFile(filePath)
	if err != nil {
		return err
	}

	if err := s.repo.SavePatientsToDB(patients); err != nil {
		return err
	}

	return nil
}

func (s *service) ClassifyPatients() error {
    patients, err := s.repo.GetAllPatients()
    if err != nil {
        return err
    }

    var trainingData []nb.Dataset
    for _, p := range patients {
        trainingData = append(trainingData, nb.Dataset{
            Diagnosis: p.Diagnosis,
            Priority:  p.Priority,
        })
    }

    probabilities := nb.TrainNaiveBayes(trainingData)

    for i := range patients {
        normalizedDiagnosis := strings.ToLower(strings.TrimSpace(patients[i].Diagnosis))
        priority := nb.PredictNaiveBayes(probabilities, normalizedDiagnosis)

        if strings.ToLower(patients[i].Category) == "ibu hamil" {
            if nb.IsSevereDiagnosis(normalizedDiagnosis) {
                priority = "tinggi"
            } else {
                priority = "sedang"
            }
        } else if patients[i].Age > 60 {
            if nb.IsSevereDiagnosis(normalizedDiagnosis) {
                priority = "tinggi"
            } else if nb.IsModerateDiagnosis(normalizedDiagnosis) {
                priority = "sedang"
            } else {
                priority = "rendah"
            }
        } else if patients[i].Age < 18 {
            if nb.IsSevereDiagnosis(normalizedDiagnosis) {
                priority = "sedang"
            } else {
                priority = "rendah"
            }
        } else {
            if nb.IsSevereDiagnosis(normalizedDiagnosis) {
                priority = "sedang"
            } else if nb.IsModerateDiagnosis(normalizedDiagnosis) {
                priority = "sedang"
            } else {
                priority = "rendah"
            }
        }

        log.Printf("Patient %s, Diagnosis: %s, Priority: %s\n", patients[i].Name, patients[i].Diagnosis, priority)
        patients[i].Priority = priority
    }

    if err := s.repo.SavePatientsToDB(patients); err != nil {
        return err
    }

    return nil
}
