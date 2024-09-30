package handler

import (
	"net/http"
	"qhealth/domain"
	"qhealth/features/doctor"
	"qhealth/helpers"
	"qhealth/helpers/middleware"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	service doctor.Service
}

func NewDoctorHandler(service doctor.Service) doctor.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Login(e echo.Context) error {
	doctorReq := domain.DoctorLogin{}
	err := e.Bind(&doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	token, err := h.service.Login(doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully login", echo.Map{
		"token": token,
	}))
}

func (h *handler) Register(e echo.Context) error {
	doctorReq := domain.DoctorRegister{}
	err := e.Bind(&doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	_, err = govalidator.ValidateStruct(doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	uploadStr, err := e.FormFile("upload_str")
	if err != nil && err !=http.ErrMissingFile {
		return helpers.CustomErr(e, "error handling image file: " + err.Error())
	}

	uploadSip, err := e.FormFile("upload_sip")
	if err != nil && err !=http.ErrMissingFile {
		return helpers.CustomErr(e, "error handling image file: " + err.Error())
	}

	uploadSertifikasi, err := e.FormFile("upload_sertifikasi")
	if err != nil && err !=http.ErrMissingFile {
		return helpers.CustomErr(e, "error handling image file: " + err.Error())
	}

	client := helpers.ConfigCloud()
	uploadStrUrl := helpers.UploadFile(uploadStr, client)
	uploadSipUrl := helpers.UploadFile(uploadSip, client)
	uploadSertifUrl := helpers.UploadFile(uploadSertifikasi, client)
	doctorReq.UploadStr = uploadStrUrl
	doctorReq.UploadSip = uploadSipUrl
	doctorReq.UploadSertifikasi = uploadSertifUrl

	err = h.service.Register(doctorReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusCreated, helpers.SuccessResponse("registered successfully", nil))
}

func (h *handler) ChangePass(e echo.Context) error {
	_, email, err := middleware.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	reqPass := domain.ReqChangePassDoctor{}
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
	_, email, err := middleware.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	newPass := domain.ChangePasswordDoctor{}
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
	email := domain.DoctorEmail{}
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

func (h *handler) DeleteProfile(e echo.Context) error {
	_, email, err := middleware.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.service.DeleteProfile(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, "User not found")
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully deleted user", nil))
}

func (h *handler) UpdateProfile(e echo.Context) error {
    _, email, err := middleware.ExtractToken(e)
    if err != nil {
        return helpers.CustomErr(e, "Invalid token")
    }

    doctorReq := domain.DoctorReq{}
    if err := e.Bind(&doctorReq); err != nil {
        return helpers.CustomErr(e, err.Error())
    }

	fileImage, err := e.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return helpers.CustomErr(e, "error handling image file: " + err.Error())
	}

	fileImageKtp, err := e.FormFile("image_ktp")
	if err != nil && err != http.ErrMissingFile { 
        return helpers.CustomErr(e, "error handling image_ktp file: " + err.Error())
    }

	client := helpers.ConfigCloud()
	imageUrl := helpers.UploadFile(fileImage, client)
	imageKtpUrl := helpers.UploadFile(fileImageKtp, client)
	doctorReq.Image = imageUrl
	doctorReq.ImageKtp = imageKtpUrl

    err = h.service.UpdateProfile(email, doctorReq)  
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data", nil))
}