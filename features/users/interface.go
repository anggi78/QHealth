package users

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateUser(user domain.User) error
		FindByEmail(email string) (domain.User, error)
		FindCodeByEmail(email string) (string, error)
		UpdatePass(email, newPass string) error
		DeleteUser(email string) error
		UpdateUser(email string, data map[string]interface{}) error
	}

	Service interface {
		Register(userReq domain.UserRegister) error
		Login(userReq domain.UserLogin) (string, error)
		ChangePass(email string, reqPass domain.ReqChangePass) error
		ChangePassForgot(email, newPass string) error
		ForgotPassword(email string) error
		DeleteUser(email string) error
		UpdateUser(email string, user domain.UserReq) error
	}

	Handler interface {
		Login(e echo.Context) error
		Register(e echo.Context) error
		ChangePass(e echo.Context) error
		ChangePassForgot(e echo.Context) error
		ForgotPassword(e echo.Context) error
		DeleteUser(e echo.Context) error
		UpdateUser(e echo.Context) error
	}
)