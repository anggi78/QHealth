package handler

import (
	"log"
	"net/http"
	"os"
	"qhealth/domain"
	"qhealth/features/users"
	"qhealth/helpers"
	"qhealth/helpers/middleware"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (h *handler) RegisterAdmin(e echo.Context) error {
    userReq := domain.UserRegister{}
    err := e.Bind(&userReq)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    if isValidAdminEmail(userReq.Email) {
        err = h.service.RegisterAdmin(userReq)
    } else {
        err = h.service.Register(userReq)
    }

    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusCreated, helpers.SuccessResponse("registered successfully", nil))
}

func isValidAdminEmail(email string) bool {
    adminEmail := os.Getenv("ADMIN_EMAIL")
    if adminEmail == "" {
        log.Fatal("ADMIN_EMAIL is not set")
    }
    return email == adminEmail
}

func (h *handler) ChangePass(e echo.Context) error {
	_, email, err := middleware.ExtractToken(e)
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
	_, email, err := middleware.ExtractToken(e)
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

func (h *handler) DeleteUser(e echo.Context) error {
	_, email, err := middleware.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.service.DeleteUser(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, "User not found")
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully deleted user", nil))
}

func (h *handler) UpdateUser(e echo.Context) error {
    _, email, err := middleware.ExtractToken(e)
    if err != nil {
        return helpers.CustomErr(e, "Invalid token")
    }

    userReq := domain.UserReq{}
    if err := e.Bind(&userReq); err != nil {
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
	userReq.Image = imageUrl
	userReq.ImageKtp = imageKtpUrl

    err = h.service.UpdateUser(email, userReq)  
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data", nil))
}

func (h *handler) InitializeRolesAndPermissions(e echo.Context) error {
    err := h.service.InitializeRolesAndPermission()
	if err != nil {
        return helpers.CustomErr(e, err.Error())
    }
	return e.JSON(http.StatusOK, helpers.SuccessResponse("roles and permissions initialized successfully", nil))
}