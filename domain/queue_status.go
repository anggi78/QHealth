package domain

import (
	"qhealth/helpers"

	"gorm.io/gorm"
)

type QueueStatus struct {
	Id   string `gorm:"PrimaryKey"`
	Name string `gorm:"not null"`
}

type QueueStatusReq struct {
	Name string `json:"name" valid:"required~your status is required"`
}

type QueueStatusResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *QueueStatus) BeforeCreate(tx *gorm.DB) error {
	s.Id = helpers.CreateId()
	return nil
}