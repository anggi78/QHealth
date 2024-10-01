package repository

import (
	"qhealth/domain"
	queuestatus "qhealth/features/queue_status"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) queuestatus.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateStatus(status domain.QueueStatus) error {
	err := r.db.Create(&status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllStatus() ([]domain.QueueStatus, error) {
	var status []domain.QueueStatus
	err := r.db.Find(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}