package service

import (
	"qhealth/domain"
	"qhealth/features/notification"
	"qhealth/helpers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type service struct {
	repo notification.Repository
	validate *validator.Validate
}

func NewNotificationService(repo notification.Repository, validate *validator.Validate) notification.Service {
	return &service{
		repo: repo,
		validate: validate,
	}
}

func (s *service) FindAllNotification(c echo.Context) ([]domain.Notification, error) {
	notification, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if notification == nil {
		return nil, helpers.CustomErr(c, "notification not found")
	}

	return notification, nil
}