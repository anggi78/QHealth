package service

import (
	"qhealth/domain"
	"qhealth/features/message"
)

type service struct {
	repo message.Repository
}

func NewMessageService(repo message.Repository) message.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllMessage() ([]domain.MessageResp, error) {
	message, err := s.repo.GetAllMessage()
	if err != nil {
		return nil, err
	}

	result := domain.ListMessageToResp(message)
	return result, nil
}