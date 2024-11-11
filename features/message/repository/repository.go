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
	// if senderRole == "user" {
	// 	var user domain.User
	// 	if err := r.db.First(&user, "id = ?", msg.IdUser).Error; err != nil {
	// 		return fmt.Errorf("sender user not found: %v", err)
	// 	}
	// } else if senderRole == "doctor" {
	// 	var doctor domain.Doctor
	// 	if err := r.db.First(&doctor, "id = ?", msg.IdDoctor).Error; err != nil {
	// 		return fmt.Errorf("sender doctor not found: %v", err)
	// 	}
	// }

	// if msg.IdDoctor != "" {
	// 	var doctor domain.Doctor
	// 	if err := r.db.First(&doctor, "id = ?", recipientId).Error; err != nil {
	// 		return fmt.Errorf("receiver doctor not found: %v", err)
	// 	}
	// } else {
	// 	var user domain.User
	// 	if err := r.db.First(&user, "id = ?", recipientId).Error; err != nil {
	// 		return fmt.Errorf("receiver user not found: %v", err)
	// 	}
	// }

	if err := r.db.Create(&msg).Error; err != nil {
		return err
	}

	var userExists bool
	if err := r.db.Table("users").Select("count(*) > 0").Where("id = ?", recipientId).Find(&userExists).Error; err != nil || !userExists {
		return fmt.Errorf("recipient ID %s tidak ditemukan di tabel users", recipientId)
	}

	newRecipientId := helpers.CreateId()
	recipientQuery := `INSERT INTO message_recipients (id, id_message, id_user, is_read) VALUES (?, ?, ?, ?)`
	if err := r.db.Exec(recipientQuery, newRecipientId, msg.Id, recipientId, false).Error; err != nil {
		return err
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
