package domain

func UserRegisterToUser(u UserRegister) User {
	return User{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
	}
}

func ReqToUser(u UserReq) User {
	return User{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
		Image:    u.Image,
		Birth:    u.Birth,
		JK:       u.JK,
		Nik:      u.Nik,
		ImageKtp: u.ImageKtp,
	}
}
