package service

import (
	"qhealth/domain"
	"qhealth/features/queue"
	"time"
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
	status, err := s.repo.GetQueueStatusByName("Menunggu")
	if err != nil {
		return err
	}

	queue := domain.ReqToQueue(queueReq)
	queue.IdQueueStatus = status.Id

	nextQueueNumber, err := s.repo.GetNextQueueNumber()
	if err != nil {
		return err
	}
	queue.QueueNumber = nextQueueNumber

	queuePosition, err := s.repo.GetQueuePosition(queue.IdDoctor, nextQueueNumber)
	if err != nil {
		return err
	}
	queue.QueuePosition = queuePosition

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

func (s *service) CallPatient(queueNumber, doctorID string) error {
	status, err := s.repo.GetQueueStatusByName("Dipanggil")
	if err != nil {
		return err
	}

	calledAt := time.Now()
	err = s.repo.UpdateQueueStatus(queueNumber, status.Id, calledAt)
	if err != nil {
		return err
	}

	return nil
}