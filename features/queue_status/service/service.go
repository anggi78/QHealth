package service

import (
	"qhealth/domain"
	queuestatus "qhealth/features/queue_status"
)

type service struct {
	repo queuestatus.Repository
}

func NewStatusService(repo queuestatus.Repository) queuestatus.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateStatus(statusReq domain.QueueStatusReq) error {
	status := domain.ReqToQueueStatus(statusReq)

	err := s.repo.CreateStatus(status)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllStatus() ([]domain.QueueStatusResp, error) {
	status, err := s.repo.GetAllStatus()

	if err != nil {
		return nil, err
	}

	result := domain.ListQueueStatusToResp(status)
	return result, nil
}