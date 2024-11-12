package repository

import (
	"fmt"
	"qhealth/domain"
	"qhealth/features/message"
	"qhealth/helpers"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) message.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveMessage(msg domain.Message, recipientId string) error {
	if err := r.db.Create(&msg).Error; err != nil {
		return err
	}

	var userExists bool
	var doctorExists bool

	if err := r.db.Table("users").Select("count(*) > 0").Where("id = ?", recipientId).Find(&userExists).Error; err != nil {
		return err
	}

	if !userExists {
		if err := r.db.Table("doctors").Select("count(*) > 0").Where("id = ?", recipientId).Find(&doctorExists).Error; err != nil {
			return err
		}
	}

	if !userExists && !doctorExists {
		return fmt.Errorf("recipient ID %s tidak ditemukan di tabel users atau doctors", recipientId)
	}

	newRecipientId := helpers.CreateId()
	var recipientQuery string

	if userExists {
		recipientQuery = `INSERT INTO message_recipients (id, id_message, id_user, is_read) VALUES (?, ?, ?, ?)`
		if err := r.db.Exec(recipientQuery, newRecipientId, msg.Id, recipientId, false).Error; err != nil {
			return err
		}
	} else if doctorExists {
		recipientQuery = `INSERT INTO message_recipients (id, id_message, id_doctor, is_read) VALUES (?, ?, ?, ?)`
		if err := r.db.Exec(recipientQuery, newRecipientId, msg.Id, recipientId, false).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) GetUnreadMessages(userID string) ([]domain.Message, error) {
	var unreadMessages []domain.Message
	err := r.db.Table("messages").
		Joins("JOIN message_recipients ON messages.id = message_recipients.id_message").
		Where("message_recipients.id_user = ? AND message_recipients.is_read = ?", userID, false).
		Find(&unreadMessages).Error

	return unreadMessages, err
}

func (r *repository) IsDoctor(senderId string) (bool, error) {
	var count int64
	err := r.db.Table("doctors").Where("id = ?", senderId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) GetAllMessage() ([]domain.Message, error) {
	var message []domain.Message
	err := r.db.Preload("User").Preload("Doctor").Find(&message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}