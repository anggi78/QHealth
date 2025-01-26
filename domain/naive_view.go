package domain

func PatientToResp(p Patient) PatientResp {
	return PatientResp{
		Id:        p.Id,
		Name:      p.Name,
		Age:       p.Age,
		Diagnosis: p.Diagnosis,
		Category:  p.Category,
	}
}

func ListPatientToResp(p []Patient) []PatientResp {
	result := []PatientResp{}
	for _, v := range p {
		data := PatientToResp(v)
		result = append(result, data)
	}
	return result
}