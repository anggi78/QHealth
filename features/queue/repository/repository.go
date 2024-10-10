package repository

import (
	"qhealth/domain"
	"qhealth/features/queue"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewQueueRepository(db *gorm.DB) queue.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateQueue(queue domain.Queue) error {
	err := r.db.Create(&queue).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllQueues() ([]domain.Queue, error) {
	var queues []domain.Queue
	err := r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").Find(&queues).Error
	if err != nil {
		return nil, err
	}
	return queues,nil
}

func (r *repository) GetQueueByID(id string) (*domain.Queue, error) {
	var queue domain.Queue
	err := r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").Where("id = ?", id).First(&queue).Error
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r *repository) GetQueueStatusByName(statusName string, status *domain.QueueStatus) error {
	err := r.db.Where("name = ?", statusName).First(status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateQueue(id string, queue domain.Queue) error {
	_, err := r.GetQueueByID(id)
	if err != nil {
		return err
	}

	err = r.db.Where("id = ?", id).Updates(&queue).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteQueue(id string) error {
	err := r.db.Delete(&domain.Queue{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}