package domain

func ReqToQueue(q QueueReq) Queue {
	return Queue{
		// QueueNumber:   q.QueueNumber,
		// QueuePosition: q.QueuePosition,
		IdUser:        q.IdUser,
		IdDoctor:      q.IdDoctor,
	}
}

func QueueToResp(q Queue) QueueResp {
	return QueueResp{
		Id:            q.Id,
		QueueNumber:   q.QueueNumber,
		QueuePosition: q.QueuePosition,
		IdUser:        q.IdUser,
		User: UserResp{
			Name:     q.User.Name,
			Email:    q.User.Email,
			Phone:    q.User.Phone,
			Address:  q.User.Address,
			Image:    q.User.Image,
			Birth:    q.User.Birth,
			JK:       q.User.JK,
			Nik:      q.User.Nik,
			ImageKtp: q.User.ImageKtp,
		},
		IdDoctor: q.IdDoctor,
		Doctor: DoctorRespToQueue{
			Name:         q.Doctor.Name,
			Spesialisasi: q.Doctor.Spesialisasi,
		},
		IdQueueStatus: q.IdQueueStatus,
		QueueStatus: QueueStatusRespToQueue{
			Name: q.QueueStatus.Name,
		},
	}
}

func ListQueueToResp(q []Queue) []QueueResp {
	result := []QueueResp{}
	for _, v := range q {
		data := QueueToResp(v)
		result = append(result, data)
	}
	return result
}
