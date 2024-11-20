package domain

func NotificationToResp(n Notification) NotificationResp {
	return NotificationResp{
		Id: n.Id,
		Type: n.Type,
		Message: n.Message,
		IsRead: n.IsRead,
		IdUser: n.IdUser,
		User: UserResp{
			Name:     n.User.Name,
			Email:    n.User.Email,
			Phone:    n.User.Phone,
			Address:  n.User.Address,
			Image:    n.User.Image,
			Birth:    n.User.Birth,
			JK:       n.User.JK,
			Nik:      n.User.Nik,
			ImageKtp: n.User.ImageKtp,
		},
	}
}

func ListNotificationToResp(n []Notification) []NotificationResp {
	result := []NotificationResp{}
	for _, v := range n {
		data := NotificationToResp(v)
		result = append(result, data)
	}
	return result
}