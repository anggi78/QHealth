package repository

import (
	"errors"
	"fmt"
	"log"
	"qhealth/domain"
	"qhealth/features/queue"
	"strconv"
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

func (r *repository) GetQueueStatusByName(statusName string) (*domain.QueueStatus, error) {
	var status domain.QueueStatus
	err := r.db.Where("name = ?", statusName).First(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *repository) GetNextQueue(doctorId string) (*domain.Queue, error) {
    var queue domain.Queue
    
    status, err := r.GetQueueStatusByName("Menunggu")
    if err != nil {
        return nil, err
    }

    log.Println("Queue status ID:", status.Id)  

    err = r.db.Preload("User").Preload("Doctor").Preload("QueueStatus").
        Where("id_doctor = ? AND id_queue_status = ?", doctorId, status.Id).
        Order("created_at ASC").Find(&queue).Error

    if err != nil {
        log.Println("Error fetching next queue:", err) 
        return nil, err
    }

    return &queue, nil
}

func (r *repository) GetNextQueueNumber() (string, error) {
	var queue domain.Queue

	err := r.db.Order("queue_number desc").First(&queue).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	lastNumber := 0
	if queue.QueueNumber != "" {
		numberPart := queue.QueueNumber[1:]
		lastNumber, err = strconv.Atoi(numberPart)
		if err != nil {
			return "", err
		}
	}
	nextNumber := lastNumber + 1
	
	if nextNumber%5 == 0 || (nextNumber+1)%5 == 0 {
		nextNumber += 2
	}

	nextQueueNumber := fmt.Sprintf("A%03d", nextNumber)
	return nextQueueNumber, nil
}

func (r *repository) GetQueuePosition(doctorID, userQueue string) (string, error) {
	var count int64
	
	status, err := r.GetQueueStatusByName("Menunggu")  
	if err != nil {
		return "", err
	}
	
	err = r.db.Model(&domain.Queue{}).
		Where("id_doctor = ? AND id_queue_status = ? AND called_at IS NULL AND queue_number < ?", doctorID, status.Id, userQueue).
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(count)), nil
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