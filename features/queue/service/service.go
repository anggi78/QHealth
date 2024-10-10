package service

import (
	"qhealth/domain"
	"qhealth/features/queue"
)

type service struct {
	repo queue.Repository
}

func NewQueueService(repo queue.Repository) queue.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateQueue(queueReq domain.QueueReq) error {
	var defaultStatus domain.QueueStatus
	err := s.repo.GetQueueStatusByName("Menunggu", &defaultStatus)
	if err != nil {
		return err
	}

	queue := domain.ReqToQueue(queueReq)

	queue.IdQueueStatus = defaultStatus.Id

	err = s.repo.CreateQueue(queue)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllQueues() ([]domain.QueueResp, error) {
	queue, err := s.repo.GetAllQueues()
	if err != nil {
		return nil, err
	}

	result := domain.ListQueueToResp(queue)
	return result, nil
}

func (s *service) GetQueueByID(id string) (*domain.QueueResp, error) {
	queue, err := s.repo.GetQueueByID(id)
	if err != nil {
		return nil, err
	}

	queueResp := domain.QueueToResp(*queue)

	return &queueResp, nil
}

func (s *service) UpdateQueue(id string, queue domain.QueueReq) error {
	data := domain.ReqToQueue(queue)

	err := s.repo.UpdateQueue(id, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteQueue(id string) error {
	return s.repo.DeleteQueue(id)
}