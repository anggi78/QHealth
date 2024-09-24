package helpers

import (
	"errors"
	"time"
)

func ParsedDate(date string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", errors.New("invalid format date")
	}
	strDate := parsedDate.Format("2006-01-02")
	return strDate, nil

}