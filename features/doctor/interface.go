package doctor

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateDoctor(doctor domain.Doctor) error
		FindByEmail(email string) (domain.Doctor, error)
		UpdatePass(email, newPass string) error
		DeleteProfile(email string) error
		UpdateProfile(email string, doctor domain.Doctor) error
	}

	Service interface {
		Register(doctorReq domain.DoctorRegister) error
		Login(doctorReq domain.DoctorLogin) (string, error)
		ChangePass(email string, reqPass domain.ReqChangePassDoctor) error 
		ChangePassForgot(email, newPass string) error
		ForgotPassword(email string) error
		DeleteProfile(email string) error
		UpdateProfile(email string, doctor domain.DoctorReq) error
	}

	Handler interface {
		Login(e echo.Context) error
		Register(e echo.Context) error
		ChangePass(e echo.Context) error
		ChangePassForgot(e echo.Context) error
		ForgotPassword(e echo.Context) error
		DeleteProfile(e echo.Context) error
		UpdateProfile(e echo.Context) error
	}
)