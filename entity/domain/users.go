package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"PrimaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string
	Phone     string
	Address   string
	Image     string
	Birth     time.Time
	JK        string
	Nik       string
	ImageKtp  string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"required~your email is required, email~invalid email format"`
	Password string `json:"password"`
}

type UserEmail struct {
	Email string `json:"email" valid:"required~your email is required, email~invalid email format"`
}

type UserRegister struct {
	Name     string `json:"user_name" valid:"required~your username is required"`
	Email    string `json:"email" valid:"required~your email is required, email~invalid email format"`
	Password string `json:"password" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Phone    string `json:"phone"`
}

type SendOtp struct {
	Code  string `json:"code" valid:"required~your code is required"`
	Email string `json:"email" valid:"required~your email is required, email~invalid email format"`
}

type ChangePassword struct {
	Password        string `json:"password" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	ConfirmPassword string `json:"confirm_password" valid:"required~your confirm password  is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type ReqChangePass struct {
	OldPass     string `json:"old_pass" valid:"required~your old password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	NewPass     string `json:"new_pass" valid:"required~your new password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	ConfirmPass string `json:"confirm_pass" valid:"required~your confirm password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Id = helpers.CreateId()
	u.Password, _ = helpers.HassPass(u.Password)
	return nil
}
