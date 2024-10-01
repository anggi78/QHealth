package domain

func ReqToQueueStatus(s QueueStatusReq) QueueStatus {
	return QueueStatus{
		Name: s.Name,
	}
}

func QueueStatusToResp(s QueueStatus) QueueStatusResp {
	return QueueStatusResp{
		Id: s.Id,
		Name: s.Name,
	}
}

func ListQueueStatusToResp(s []QueueStatus) []QueueStatusResp {
	result := []QueueStatusResp{}
	for _, v := range s {
		data := QueueStatusToResp(v)
		result = append(result, data)
	}
	return result
}