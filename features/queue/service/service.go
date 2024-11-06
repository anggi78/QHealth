package service

import (
	"fmt"
	"qhealth/domain"
	"qhealth/features/queue"
	"strconv"
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

	lastQueue, err := s.repo.GetLastQueue()
	if err != nil {
		return err
	}

	lastNumber := 0
	if lastQueue.QueueNumber != "" {
		numberPart := lastQueue.QueueNumber[1:]
		lastNumber, err = strconv.Atoi(numberPart)
		if err != nil {
			return err
		}
	}
	nextNumber := lastNumber + 1

	if nextNumber%5 == 0 || (nextNumber+1)%5 == 0 {
		for i := 0; i < 2; i++ {
			offlineNumber := fmt.Sprintf("A%03d", nextNumber+i)
			offlineQueuePosition := strconv.Itoa(lastNumber + i + 1)
			err = s.repo.CreateOfflineQueue(offlineNumber, offlineQueuePosition, status.Id)
			if err != nil {
				return err
			}
		}
		nextNumber += 2
	}

	nextQueueNumber := fmt.Sprintf("A%03d", nextNumber)
	queue.QueueNumber = nextQueueNumber

	doctorID := ""
	if queue.IdDoctor != nil {
		doctorID = *queue.IdDoctor
	}

	count, err := s.repo.CountWaitingQueues(doctorID, nextQueueNumber, status.Id)
	if err != nil {
		return err
	}

	queue.QueuePosition = strconv.Itoa(int(count))

	return s.repo.CreateQueue(queue)
}

func (s *service) GetAllQueues() ([]domain.QueueResp, error) {
	queues, err := s.repo.GetAllQueues()
	if err != nil {
		return nil, err
	}

	activeWaitingIndex := 0
	for i := range queues {
		switch queues[i].QueueStatus.Name {
		case "Menunggu":
			queues[i].QueuePosition = strconv.Itoa(activeWaitingIndex)
			activeWaitingIndex++
		case "Dipanggil", "Selesai", "Dibatalkan":
			queues[i].QueuePosition = "0"
		}
	}

	result := domain.ListQueueToResp(queues)
	return result, nil
}

func (s *service) GetAllQueuesAdmin(admin bool) ([]domain.QueueResp, error) {
	queues, err := s.repo.GetAllQueues()
	if err != nil {
		return nil, err
	}

	for i := range queues {
		if admin && queues[i].QueueStatus.Name == "Menunggu" {
			queues[i].QueueStatus.Name = "Terkonfirmasi"
		}
	}

	activeWaitingIndex := 0
	for i := range queues {
		switch queues[i].QueueStatus.Name {
		case "Terkonfirmasi":
			queues[i].QueuePosition = strconv.Itoa(activeWaitingIndex)
			activeWaitingIndex++
		case "Dipanggil", "Selesai", "Dibatalkan":
			queues[i].QueuePosition = "0"
		}
	}

	result := domain.ListQueueToResp(queues)
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

func (s *service) CompleteQueue(queueNumber, doctorID string) error {
	status, err := s.repo.GetQueueStatusByName("Selesai")
	if err != nil {
		return err
	}

	completedAt := time.Now()
	err = s.repo.UpdateQueueStatus(queueNumber, status.Id, completedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CancelQueue(queueNumber, doctorID string) error {
	status, err := s.repo.GetQueueStatusByName("Dibatalkan")
	if err != nil {
		return err
	}

	emptyTime := time.Time{}
	err = s.repo.UpdateQueueStatus(queueNumber, status.Id, emptyTime)
	if err != nil {
		return err
	}

	return nil
}
