package handler

import (
	"fmt"
	"net/http"
	"qhealth/domain"
	"qhealth/features/users"
	"qhealth/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service users.Service
}

func NewUserHandler(service users.Service) users.Handler {
	return &handler{service: service}
}

func (h *handler) Login(e echo.Context) error {
	userReq := domain.UserLogin{}
	err := e.Bind(&userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	token, err := h.service.Login(userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully login", echo.Map{
		"token": token,
	}))
}

func (h *handler) Register(e echo.Context) error {
	userReq := domain.UserRegister{}
	err := e.Bind(&userReq)
	if err != nil {
		fmt.Println("err", err)
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.service.Register(userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusCreated, helpers.SuccessResponse("registered successfully", nil))
}

func (h *handler) ChangePass(e echo.Context) error {
	_, email, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	reqPass := domain.ReqChangePass{}
	err = e.Bind(&reqPass)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(reqPass)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.service.ChangePass(email, reqPass)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully change password", nil))
}

func (h *handler) ChangePassForgot(e echo.Context) error {
	_, email, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	newPass := domain.ChangePassword{}
	err = e.Bind(&newPass)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(newPass)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	if newPass.Password != newPass.ConfirmPassword {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "password does not match",
		})
	}

	err = h.service.ChangePassForgot(email, newPass.Password)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}
	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully change password", nil))
}

func (h *handler) ForgotPassword(e echo.Context) error {
	email := domain.UserEmail{}
	err := e.Bind(&email)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(email)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.service.ForgotPassword(email.Email)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully send an email", nil))
}