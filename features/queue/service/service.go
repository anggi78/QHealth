package service

import (
	"fmt"
	"qhealth/domain"
	"qhealth/features/queue"
	"qhealth/helpers"
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

	queueType := ""
	queuePrefix := ""
	if (queueReq.Age >= 16 && queueReq.Age <= 59) || queueReq.IsHajjCheck ||
		queueReq.IsDentalPatient || queueReq.IsTBTreatment || queueReq.IsHospitalReferral {
		queueType = "Umum"
		queuePrefix = "A"
	} else if queueReq.Age <= 15 || queueReq.Age >= 60 ||
		queueReq.IsDoctorCertificate || queueReq.IsPregnantReferral {
		queueType = "Khusus"
		queuePrefix = "B"
	} else {
		return fmt.Errorf("kategori antrian tidak valid")
	}
	queue.QueueType = queueType

	lastQueue, err := s.repo.GetLastQueue(queueType)
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
			offlineNumber := fmt.Sprintf("%s%03d", queuePrefix, nextNumber+i)
			offlineQueuePosition := strconv.Itoa(lastNumber + i + 1)
			err = s.repo.CreateOfflineQueue(offlineNumber, offlineQueuePosition, status.Id)
			if err != nil {
				return err
			}
		}
		nextNumber += 2
	}

	nextQueueNumber := fmt.Sprintf("%s%03d", queuePrefix, nextNumber)
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

func (s *service) GetAllQueues(page, pageSize int) ([]domain.QueueResp, int, error) {
	offset := (page - 1) * pageSize

	queues, totalItems, err := s.repo.GetAllQueues(offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	if len(queues) == 0 {
		result := domain.ListQueueToResp(queues)
		return result, totalItems, nil
	}

	totalWaitingBefore, err := s.repo.CountWaitingQueuesBeforePage(queues[0].QueueNumber, "")
	if err != nil {
		return nil, 0, err
	}

	activeWaitingIndex := int(totalWaitingBefore) + 0

	for i := range queues {
		switch queues[i].QueueStatus.Name {
		case "Menunggu":
			oldPosition, err := strconv.Atoi(queues[i].QueuePosition)
			if err != nil {
				fmt.Printf("Invalid queue position for queue ID %s: %v\n", queues[i].Id, err)
				continue
			}

			queues[i].QueuePosition = strconv.Itoa(activeWaitingIndex)

			if oldPosition > activeWaitingIndex {
				err := helpers.SendQueueNotification(queues[i].User.Email)
				if err != nil {
					fmt.Printf("Failed to send notification for queue ID %s: %v\n", queues[i].Id, err)
				}
			}

			activeWaitingIndex++
		case "Dipanggil", "Selesai", "Dibatalkan":
			queues[i].QueuePosition = "0"
		}
	}
	fmt.Println("Total Waiting Before:", totalWaitingBefore)

	fmt.Println("First Queue Number on Page:", queues[0].QueueNumber)
	result := domain.ListQueueToResp(queues)
	return result, totalItems, nil
}

func (s *service) GetAllQueuesAdmin(admin bool, page, pageSize int) ([]domain.QueueResp, int, error) {
	queues, _, err := s.repo.GetAllQueues(page, pageSize)
	if err != nil {
		return nil, 0, err
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
	return result, 0, nil
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
