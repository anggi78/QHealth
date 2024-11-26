package main

import (
	//"log"
	"net/http"
	configs "qhealth/app/drivers"
	"qhealth/app/routes"
	message "qhealth/features/message/repository"
	users "qhealth/features/users/repository"
	doctor "qhealth/features/doctor/repository"
	notification "qhealth/features/notification/repository"
	"qhealth/features/message/ws"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db := configs.InitDB()
	messageRepo := message.NewMessageRepository(db)
    userRepo := users.NewUserRepository(db)
	doctorRepo := doctor.NewDoctorRepository(db)
	notifRepo := notification.NewNotificationRepository(db)
	hub := ws.NewHub(messageRepo, userRepo, doctorRepo, notifRepo)
	go hub.Run()
	validate := validator.New()

	routes.Routes(e, db, hub, validate)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"ngrok-skip-browser-warning",
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	e.Logger.Fatal(e.Start(":8089"))
}
