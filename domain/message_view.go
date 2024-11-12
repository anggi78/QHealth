package domain

func MessageToResp(m Message) MessageResp {
	return MessageResp{
		Id:              m.Id,
		MessageBody:     m.MessageBody,
		CreateDate:      m.CreateDate,
		IdParentMessage: m.IdParentMessage,
		ParentMessage:   m.ParentMessage,
		IdUser:          m.IdUser,
		User: UserResp{
			Name:     m.User.Name,
			Email:    m.User.Email,
			Phone:    m.User.Phone,
			Address:  m.User.Address,
			Image:    m.User.Image,
			Birth:    m.User.Birth,
			JK:       m.User.JK,
			Nik:      m.User.Nik,
			ImageKtp: m.User.ImageKtp,
		},
		IdDoctor: m.IdDoctor,
		Doctor: DoctorRespToQueue{
			Name:         m.Doctor.Name,
			Spesialisasi: m.Doctor.Spesialisasi,
		},
	}
}

func ListMessageToResp(m []Message) []MessageResp {
	result := []MessageResp{}
	for _, v := range m {
		data := MessageToResp(v)
		result = append(result, data)
	}
	return result
}