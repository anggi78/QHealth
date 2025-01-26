package naivebayes

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		SavePatientsToDB(patients []domain.Patient) error
		GetAllPatients() ([]domain.Patient, error)
	}

	Service interface {
		GetAllPatients() ([]domain.PatientResp, error)
		ImportPatientsFromExcel(filePath string) error
	}

	Handler interface {
		GetAllPatients(e echo.Context) error
	}
)