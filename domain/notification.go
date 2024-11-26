package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	Id            string         `gorm:"primaryKey" json:"id"`
	Type          string         `json:"type"`
	Message       string         `json:"message"`
	IsRead        bool           `json:"is_read"`
	RecipientType string         `json:"recipient_type"`
	RecipientId   string         `json:"recipient_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
type NotificationResp struct {
	Id            string `gorm:"primaryKey" json:"id"`
	Type          string `json:"type"`
	Message       string `json:"message"`
	IsRead        bool   `json:"id_read"`
	RecipientType string `json:"recipient_type"`
	RecipientId   string `json:"recipient_id"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	n.Id = helpers.CreateId()
	return nil
}
