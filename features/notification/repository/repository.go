package repository

import (
	"qhealth/domain"
	"qhealth/features/notification"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) notification.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]domain.Notification, error) {
	var notification []domain.Notification

	result := r.db.Where("deleted_at IS NULL").Preload("User").Order("created_at DESC").Find(&notification)
	if result.Error != nil {
		return nil, result.Error
	}

	return notification, nil
}

func (r *repository) SaveNotification(notification domain.Notification) error {
	result := r.db.Create(&notification)
	return result.Error
}