package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id              string         `gorm:"primaryKey" json:"id"`
	MessageBody     string         `json:"body"`
	CreateDate      time.Time      `json:"create_date"`
	IdParentMessage *string        `json:"id_parent_message"`
	IdUser          string         `json:"id_user"`
	IdDoctor        string         `json:"id_doctor"`
	User            User           `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Doctor          Doctor         `gorm:"foreignKey:IdDoctor;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"doctor"`
	ParentMessage   *Message       `gorm:"foreignKey:IdParentMessage;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent_message"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type MessageResp struct {
	Id              string            `json:"id"`
	MessageBody     string            `json:"body"`
	CreateDate      time.Time         `json:"create_date"`
	IdParentMessage *string           `json:"id_parent_message"`
	ParentMessage   *Message          `json:"parent_message"`
	IdUser          string            `json:"id_user"`
	User            UserResp          `json:"user"`
	IdDoctor        string            `json:"id_doctor"`
	Doctor          DoctorRespToQueue `json:"doctor"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.Id = helpers.CreateId()
	return nil
}
