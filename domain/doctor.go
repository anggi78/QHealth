package domain

import (
	"qhealth/helpers"
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	Id                string `gorm:"PrimaryKey"`
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null;unique"`
	Password          string
	Phone             string
	Address           string
	Image             string
	Birth             *string
	JK                string
	Nik               string
	ImageKtp          string
	Spesialisasi      string
	Experience        string
	NumberStr         string
	NumberSip         string
	Education         string
	UploadStr         string
	UploadSip         string
	UploadSertifikasi string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type DoctorResp struct {
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

type DoctorRespToQueue struct {
	Name string  `json:"name"`
	Spesialisasi string  `json:"spesialisasi"`
}

type DoctorLogin struct {
	Email    string `json:"email" valid:"required~your email is required, email~invalid email format"`
	Password string `json:"password"`
}

type DoctorEmail struct {
	Email string `json:"email" valid:"required~your email is required, email~invalid email format"`
}

type DoctorRegister struct {
	Name              string `json:"name" form:"name" valid:"required~your name is required"`
	Email             string `json:"email" form:"email" valid:"required~your email is required, email~invalid email format"`
	Password          string `json:"password" form:"password" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Phone             string `json:"phone" form:"phone"`
	Spesialisasi      string `json:"spesialisasi" form:"spesialisasi"`
	Experience        string `json:"experience" form:"experience"`
	NumberStr         string `json:"number_str" form:"number_str"`
	NumberSip         string `json:"number_sip" form:"number_sip"`
	Education         string `json:"education" form:"education"`
	UploadStr         string `json:"upload_str" form:"upload_str"`
	UploadSip         string `json:"upload_sip" form:"upload_sip"`
	UploadSertifikasi string `json:"upload_sertifikasi" form:"upload_sertifikasi"`
}

type DoctorReq struct {
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

type ChangePasswordDoctor struct {
	Password        string `json:"password" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	ConfirmPassword string `json:"confirm_password" valid:"required~your confirm password  is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type ReqChangePassDoctor struct {
	OldPass     string `json:"old_pass" valid:"required~your old password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	NewPass     string `json:"new_pass" valid:"required~your new password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	ConfirmPass string `json:"confirm_pass" valid:"required~your confirm password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (d *Doctor) BeforeCreate(tx *gorm.DB) error {
	d.Id = helpers.CreateId()
	d.Password, _ = helpers.HassPass(d.Password)
	return nil
}