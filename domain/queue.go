package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Queue struct {
	Id            string `gorm:"PrimaryKey"`
	QueueNumber   string
	QueuePosition string
	IdUser        string
	IdDoctor      string
	IdQueueStatus string
	CalledAt      time.Time
	User          User           `gorm:"foreignKey:IdUser;references:Id"`
	Doctor        Doctor         `gorm:"foreignKey:IdDoctor;references:Id"`
	QueueStatus   QueueStatus    `gorm:"foreignKey:IdQueueStatus;references:Id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type QueueReq struct {
	QueueNumber   string `json:"queue_number"`
	QueuePosition string `json:"queue_position"`
	IdUser        string
	IdDoctor      string `json:"id_doctor"`
}

type QueueResp struct {
	Id            string                 `json:"id"`
	QueueNumber   string                 `json:"queue_number"`
	QueuePosition string                 `json:"queue_position"`
	IdUser        string                 `json:"id_user"`
	User          UserResp               `json:"user"`
	IdDoctor      string                 `json:"id_doctor"`
	Doctor        DoctorRespToQueue      `json:"doctor"`
	IdQueueStatus string                 `json:"id_queue_status"`
	QueueStatus   QueueStatusRespToQueue `json:"queue_status"`
	CalledAt      string                 `json:"called_at"`
}

func (q *Queue) BeforeCreate(tx *gorm.DB) error {
	q.Id = helpers.CreateId()
	return nil
}
