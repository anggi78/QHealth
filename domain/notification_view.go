package domain

func NotificationToResp(n Notification) NotificationResp {
	return NotificationResp{
		Id:            n.Id,
		Type:          n.Type,
		Message:       n.Message,
		IsRead:        n.IsRead,
		RecipientType: n.RecipientType,
		RecipientId:   n.RecipientId,
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
