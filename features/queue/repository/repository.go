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
		QueueNumber: queueNumber,
		QueuePosition: queuePosition,
		IdQueueStatus: statusId,
		IdDoctor: nil,
		IdUser: nil,
	}
	err := r.db.Create(&queue).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllQueues() ([]domain.Queue, error) {
    var queues []domain.Queue
    err := r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").
        Order("queue_number").Find(&queues).Error
    if err != nil {
        return nil, err
    }

    return queues, nil
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

func (r *repository) GetLastQueue() (domain.Queue, error) {
	var queue domain.Queue
	err := r.db.Order("queue_number desc").First(&queue).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Queue{}, err
	}
	return queue, nil
}

func (r *repository) CountWaitingQueues(doctorID, userQueue string, statusID string) (int64, error) {
    var count int64
    err := r.db.Model(&domain.Queue{}).
        Where("id_doctor = ? AND id_queue_status = ? AND called_at IS NULL AND queue_number < ?", doctorID, statusID, userQueue).
        Count(&count).Error
    return count, err
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

func (r *repository) UpdateQueueStatus(queueNumber, statusID string, calledAt time.Time) error {
	err := r.db.Model(&domain.Queue{}).Where("queue_number = ?", queueNumber).
			Updates(map[string]interface{}{"id_queue_status": statusID, "called_at": calledAt}).Error
	return err
}