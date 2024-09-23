package domain

func UserRegisterToUser(u UserRegister) User {
	return User{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
	}
}
