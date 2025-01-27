package naivebayes

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		SavePatientsToDB(patients []domain.Patient) error
		GetAllPatients() ([]domain.Patient, error)
		GetPatientsByPriority(priority string) ([]domain.Patient, error)
	}

	Service interface {
		GetAllPatients() ([]domain.PatientResp, error)
		GetPatientsByPriority(priority string) ([]domain.PatientResp, error)
		ImportPatientsFromExcel(filePath string) error
		ClassifyPatients() error
	}

	Handler interface {
		GetAllPatients(e echo.Context) error
		ClassifyPatients(e echo.Context) error
	}
)