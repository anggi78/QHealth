package middleware

import (
	"net/http"
	"qhealth/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Middleware struct {
    DB *gorm.DB
}

func NewMiddleware(db *gorm.DB) *Middleware {
    return &Middleware{DB: db}
}

func (m *Middleware) Authorize(permissionType string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Ambil user dari context yang sudah di-set oleh middleware JwtMiddleware
            userToken, ok := c.Get("user").(*jwt.Token)
            if !ok || !userToken.Valid {
                return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
            }

            // Ambil claims dari token
            claims, ok := userToken.Claims.(jwt.MapClaims)
            if !ok {
                return echo.NewHTTPError(http.StatusForbidden, "Invalid token claims")
            }

            userId, okId := claims["id"].(string)
            _, okEmail := claims["email"].(string)
            if !okId || !okEmail {
                return echo.NewHTTPError(http.StatusForbidden, "Invalid token data")
            }

            // Temukan user di database
            var user domain.User
            err := m.DB.Where("id = ?", userId).First(&user).Error
            if err != nil {
                return echo.NewHTTPError(http.StatusForbidden, "user not found")
            }

            // Ambil role permissions berdasarkan user role
            var rolePermissions domain.RolePermissions
            err = m.DB.Where("id_role = ?", user.IdRole).First(&rolePermissions).Error
            if err != nil {
				logrus.Printf("User ID Role: %s", user.IdRole)

                return echo.NewHTTPError(http.StatusForbidden, "role not found or permissions not set")
            }

            // Periksa permissions
            switch permissionType {
            case "read":
                if !rolePermissions.CanRead {
                    return echo.NewHTTPError(http.StatusForbidden, "access denied")
                }
            case "edit":
                if !rolePermissions.CanEdit {
                    return echo.NewHTTPError(http.StatusForbidden, "access denied")
                }
            case "delete":
                if !rolePermissions.CanDelete {
                    return echo.NewHTTPError(http.StatusForbidden, "access denied")
                }
            case "create":
                if !rolePermissions.CanCreate {
                    return echo.NewHTTPError(http.StatusForbidden, "access denied")
                }
            default:
                return echo.NewHTTPError(http.StatusForbidden, "permission not recognized")
            }

            return next(c)
        }
    }
}
