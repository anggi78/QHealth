package role

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateRole(role domain.Role) error
		GetAllRole() ([]domain.Role, error)
	}

	Service interface {
		CreateRole(roleReq domain.RoleReq) error
		GetAllRole() ([]domain.RoleResp, error)
	}

	Handler interface {
		CreateRole(e echo.Context) error
		GetAllRole(e echo.Context) error
	}
)