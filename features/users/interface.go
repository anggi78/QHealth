package users

import (
	"qhealth/entity/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateUser(user domain.User) error
		FindByEmail(email string) (domain.User, error)
		FindCodeByEmail(email string) (string, error)
		UpdatePass(email, newPass string) error
	}

	Service interface {
		Register(userReq domain.UserRegister) error
		Login(userReq domain.UserLogin) (string, error)
		ChangePass(email string, reqPass domain.ReqChangePass) error
		ChangePassForgot(email, newPass string) error
		ForgotPassword(email string) error
	}

	Handler interface {
		Login(e echo.Context) error
		Register(e echo.Context) error
		ChangePass(e echo.Context) error
		ChangePassForgot(e echo.Context) error
		ForgotPassword(e echo.Context) error
	}
)