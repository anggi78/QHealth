package repository

import (
	"errors"
	"qhealth/domain"
	"qhealth/features/queue"
	"time"

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

func (r *repository) CreateOfflineQueue(queueNumber, queuePosition, statusId string) error {
	queue := domain.Queue{
		QueueNumber:   queueNumber,
		QueuePosition: queuePosition,
		IdQueueStatus: statusId,
		IdDoctor:      nil,
		IdUser:        nil,
	}
	err := r.db.Create(&queue).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllQueues(offset, limit int) ([]domain.Queue, int, error) {
	var queues []domain.Queue
	var totalItems int64

	err := r.db.Model(&domain.Queue{}).Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").
		Order("created_at ASC").Order("queue_number").Offset(0).
		Offset(offset).Limit(limit).Find(&queues).Error
	if err != nil {
		return nil, 0, err
	}

	return queues, int(totalItems), nil
}

func (r *repository) GetQueueByID(id string) (*domain.Queue, error) {
	var queue domain.Queue
	err := r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").Where("id = ?", id).First(&queue).Error
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r *repository) GetQueueStatusByName(statusName string) (*domain.QueueStatus, error) {
	var status domain.QueueStatus
	err := r.db.Where("name = ?", statusName).First(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *repository) GetLastQueue(queueType string) (*domain.Queue, error) {
	// var queue domain.Queue
	// err := r.db.Where("queue_type = ?", queueType).Order("queue_number desc").First(&queue).Error
	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return domain.Queue{}, err
	// }
	// return queue, nil
	var queue domain.Queue
	query := r.db.Order("created_at DESC")
	if queueType != "" {
		query = query.Where("queue_type = ?", queueType)
	}

	err := query.First(&queue).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &queue, nil
}

func (r *repository) CountWaitingQueues(doctorID, userQueue, statusID string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Queue{}).
		Where("id_doctor = ? AND id_queue_status = ? AND called_at IS NULL AND queue_number < ?", doctorID, statusID, userQueue).
		Count(&count).Error
	return count, err
}

func (r *repository) CountWaitingQueuesBeforePage(queueNumber, statusID string) (int64, error) {
	var count int64

	err := r.db.Model(&domain.Queue{}).
		Where("queue_number < ?", queueNumber).
		Where("id_queue_status = ?", statusID). 
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
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

func (r *repository) UpdateQueuePosition(Id, newPosition string) error {
	err := r.db.Model(&domain.Queue{}).Where("id = ?", Id).
		Update("queue_position", newPosition).Error
	return err
}

func (r *repository) DeleteQueue(id string) error {
	err := r.db.Delete(&domain.Queue{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateQueueStatus(queueNumber, statusID string, calledAt time.Time) error {
	err := r.db.Model(&domain.Queue{}).Where("queue_number = ?", queueNumber).
		Updates(map[string]interface{}{"id_queue_status": statusID, "called_at": calledAt}).Error
	return err
}
