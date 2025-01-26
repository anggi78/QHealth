package service

import (
	"qhealth/domain"
	naivebayes "qhealth/features/naive-bayes"
	"qhealth/helpers/excel"
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
