package domain

func DoctorRegisterToDoctor(d DoctorRegister) Doctor {
	return Doctor{
		Name:              d.Name,
		Email:             d.Email,
		Password:          d.Password,
		Phone:             d.Phone,
		Spesialisasi:      d.Spesialisasi,
		Experience:        d.Experience,
		NumberStr:         d.NumberStr,
		NumberSip:         d.NumberSip,
		Education:         d.Education,
		UploadStr:         d.UploadStr,
		UploadSip:         d.UploadSip,
		UploadSertifikasi: d.UploadSertifikasi,
	}
}

func ReqToDoctor(d DoctorReq) Doctor {
	return Doctor{
		Name:     d.Name,
		Email:    d.Email,
		Phone:    d.Phone,
		Address:  d.Address,
		Image:    d.Image,
		Birth:    d.Birth,
		JK:       d.JK,
		Nik:      d.Nik,
		ImageKtp: d.ImageKtp,
	}
}
