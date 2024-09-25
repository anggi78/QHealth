package domain

import (
	"qhealth/helpers"

	"gorm.io/gorm"
)

type Role struct {
	Id   string `gorm:"PrimaryKey"`
	Name string `gorm:"not null"`
}

func (Role) TableName() string {
	return "role"
}

type RoleReq struct {
	Name string `json:"role" valid:"required~your role is required"`
}

type RoleResp struct {
	Id   string `json:"id"`
	Name string `json:"role"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.Id = helpers.CreateId()
	return nil
}
