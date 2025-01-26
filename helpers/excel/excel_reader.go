package excel

import (
	"fmt"
	"qhealth/domain"

	"github.com/xuri/excelize/v2"
)

func ReadExcelFile(filePath string) ([]domain.Patient, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

	var patients []domain.Patient
	for i, row := range rows {
		if i == 0 {
			continue 
		}
		if len(row) < 4 {
			continue 
		}

		patient := domain.Patient{
			Name:      row[0],
			Age:       parseInt(row[1]),
			Diagnosis: row[2],
			Category:  row[3],
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

func parseInt(value string) int {
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}
