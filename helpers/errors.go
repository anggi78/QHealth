package helpers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckError(e echo.Context, err, sub, message string, code int) (echo.Context, bool) {
	if strings.Contains(err, sub) {
		response := ErrorResponseJson{
			Status:  false,
			Message: message,
		}
		e.JSON(code, response)
		return e, true
	}
	return e, false
}

func CustomErr(e echo.Context, err string) error {
	e, ok := CheckError(e, err, "your name is required", "your name is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "uni_users_user_name", "username is already exist", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "duplicate key value violates unique constraint \"users_email_key\"", "email is already exist", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "Password has to have a minimum length of 6 characters", "Password has to have a minimum length of 6 characters", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "not found", "data is not found", http.StatusNotFound)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "poverty_census_nik_key", "nik is already exist ", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "poverty_census_kk_number_key", "kk number is already exist ", http.StatusBadRequest)
	if ok {
		return nil
	}

	response := ErrorResponseJson{
		Status: false,
		Message: err,
	}
	return e.JSON(http.StatusBadRequest, response)
}
