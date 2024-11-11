package router

// import (
// 	"qhealth/features/message/handler"
// 	"qhealth/features/message/ws"
// 	"qhealth/helpers/middleware"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// )

// func MessageRoute(m *echo.Group, db *gorm.DB, hub *ws.Hub) {
// 	message := m.Group("/msg", middleware.JwtMiddleware())
// 	message.GET("/ws/message", func(c echo.Context) error {
// 		return handler.MessageHandler(hub, c)
// 	})
// }