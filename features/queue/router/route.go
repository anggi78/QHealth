package router

import (
	"qhealth/features/queue/handler"
	"qhealth/features/queue/repository"
	"qhealth/features/queue/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QueueRoute(q *echo.Group, db *gorm.DB) {
	repo := repository.NewQueueRepository(db)
	serv := service.NewQueueService(repo)
	handler := handler.NewQueueService(serv)

	queue := q.Group("/queue", middleware.JwtMiddleware())
	mw := middleware.NewMiddleware(db)

	//user
	queue.POST("", handler.CreateQueue)
	queue.GET("/:id", handler.GetQueueById)
	queue.GET("", handler.GetAllQueues)

	//doctor
	queue.POST("/call/:queue_number", handler.CallNextPatient)
	queue.POST("/complete/:queue_number", handler.CompletePatient)
	queue.POST("/cancel/:queue_number", handler.CancelQueuePatient)

	//admin
	queue.GET("/admin", handler.GetAllQueuesAdmin, mw.AuthorizeAdmin())
	queue.PUT("/:id", handler.UpdateQueue, mw.AuthorizeAdmin())
	queue.DELETE("/:id", handler.DeleteQueue, mw.AuthorizeAdmin())
}