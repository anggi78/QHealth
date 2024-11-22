package users

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateUser(user domain.User) error
		FindByEmail(email string) (domain.User, error)
		FindById(id string) (domain.User, error)
		//FindCodeByEmail(email string) (string, error)
		UpdatePass(email, newPass string) error
		DeleteUser(email string) error
		UpdateUser(email string, user domain.User) error
		GetRoleByName(name string) (domain.Role, error)
		FindRoleByName(name string, role *domain.Role) error
		FindRolePermissionByRoleId(roleId string, permission *domain.RolePermissions) error
		CreateRole(role *domain.Role) error
		CreateRolePermission(permission *domain.RolePermissions) error
		UpdateRolePermission(permission *domain.RolePermissions) error
	}

	Service interface {
		Register(userReq domain.UserRegister) error
		RegisterAdmin(adminReq domain.UserRegister) error
		Login(userReq domain.UserLogin) (string, error)
		ChangePass(email string, reqPass domain.ReqChangePass) error
		ChangePassForgot(email, newPass string) error
		ForgotPassword(email string) error
		DeleteUser(email string) error
		UpdateUser(email string, user domain.UserReq) error
		InitializeRolesAndPermission() error
	}

	Handler interface {
		Login(e echo.Context) error
		Register(e echo.Context) error
		RegisterAdmin(e echo.Context) error
		ChangePass(e echo.Context) error
		ChangePassForgot(e echo.Context) error
		ForgotPassword(e echo.Context) error
		DeleteUser(e echo.Context) error
		UpdateUser(e echo.Context) error
		InitializeRolesAndPermissions(c echo.Context) error
	}
)