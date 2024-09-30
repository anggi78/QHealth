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
	Birth     *string
	JK        string
	Nik       string
	ImageKtp  string
	IdRole    string
	Role      Role `gorm:"foreignKey:IdRole"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResp struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
	Image    string  `json:"image"`
	Birth    *string `json:"birth"`
	JK       string  `json:"jk"`
	Nik      string  `json:"nik"`
	ImageKtp string  `json:"image_ktp"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"required~your email is required, email~invalid email format"`
	Password string `json:"password"`
}

type UserEmail struct {
	Email string `json:"email" valid:"required~your email is required, email~invalid email format"`
}

type UserRegister struct {
	Name     string `json:"name" valid:"required~your name is required"`
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

type UserReq struct {
	Name     string  `json:"name" form:"name" valid:"required~your username is required"`
	Email    string  `json:"email" form:"email" valid:"required~your email is required, email~invalid email format"`
	Phone    string  `json:"phone" form:"phone" valid:"required~your phone is required"`
	Address  string  `json:"address" form:"address" valid:"required~your address is required"`
	Image    string  `json:"image" form:"image" valid:"required~your image is required"`
	Birth    *string `json:"birth" form:"birth" valid:"required~your birth is required"`
	JK       string  `json:"jk" form:"jk" valid:"required~your jk is required"`
	Nik      string  `json:"nik" form:"nik" valid:"required~your nik is required"`
	ImageKtp string  `json:"image_ktp" form:"image_ktp" valid:"required~your image_ktp is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Id = helpers.CreateId()
	u.Password, _ = helpers.HassPass(u.Password)
	return nil
}
