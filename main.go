package main

import (
	"log"
	"net/http"
	configs "qhealth/app/drivers"
	"qhealth/app/routes"
	doctor "qhealth/features/doctor/repository"
	message "qhealth/features/message/repository"
	"qhealth/features/message/ws"
	naivebayes "qhealth/features/naive-bayes/repository"
	naive "qhealth/features/naive-bayes/service"
	notification "qhealth/features/notification/repository"
	users "qhealth/features/users/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db := configs.InitDB()

	repo := naivebayes.NewNaiveRepository(db)
	service := naive.NewNaiveService(repo)

	filePath := "../diagnosis.xlsx"

	if err := service.ImportPatientsFromExcel(filePath); err != nil {
		log.Fatalf("Failed to import patients: %v", err)
	} else {
		log.Println("Data successfully imported to the database")
	}

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
