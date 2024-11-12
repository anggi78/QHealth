package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type MessageRecipient struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	IsRead    bool           `json:"is_read"`
	IdMessage string         `json:"id_message"`
	IdUser    string         `json:"id_user"`
	IdDoctor  string         `json:"id_doctor"`
	User      User           `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Doctor    Doctor         `gorm:"foreignKey:IdDoctor;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"doctor"`
	Message   Message        `gorm:"foreignKey:IdMessage;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ms *MessageRecipient) BeforeCreate(tx *gorm.DB) error {
	ms.Id = helpers.CreateId()
	return nil
}
