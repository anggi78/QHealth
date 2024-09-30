package middleware

import (
	"errors"
	"net/http"
	configs "qhealth/app/drivers"
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

func GetPermissionByEmail(email string) (domain.RolePermissions, error) {
	var user domain.User
	var rolePermission domain.RolePermissions
	db := configs.InitDB()

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return rolePermission, errors.New("user not found")
	}

	if err := db.Where("id_role = ?", user.IdRole).First(&rolePermission).Error; err != nil {
		return rolePermission, errors.New("permission not found")
	}

	return rolePermission, nil
}

func CheckPermission(action string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, email, err := ExtractToken(c)
			if err != nil {
				return c.JSON(http.StatusForbidden, "invalid token")
			}

			permissions, err := GetPermissionByEmail(email)

			switch action {
			case "read":
				if !permissions.CanRead {
					return c.JSON(http.StatusForbidden, "you do not have permission to read")
                }
            case "edit":
                if !permissions.CanEdit {
                    return c.JSON(http.StatusForbidden, "you do not have permission to edit")
                }
            case "delete":
                if !permissions.CanDelete {
                    return c.JSON(http.StatusForbidden, "you do not have permission to delete")
                }
            case "create":
                if !permissions.CanCreate {
                    return c.JSON(http.StatusForbidden, "you do not have permission to create")
                }
            default:
                return c.JSON(http.StatusForbidden, "invalid action")
			}
			return next(c)
		}
	}
}