package domain

import (
	"qhealth/helpers"

	"gorm.io/gorm"
)

type RolePermissions struct {
	Id        string `gorm:"primaryKey"`
	CanCreate bool   `gorm:"default:false"`
	CanRead   bool   `gorm:"default:false"`
	CanEdit   bool   `gorm:"default:false"`
	CanDelete bool   `gorm:"default:false"`
	IdRole    string 
	Role      Role   `gorm:"foreignKey:IdRole;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RolePermissionResp struct {
	Id        string `json:"id"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanEdit   bool   `json:"can_edit"`
	CanDelete bool   `json:"can_delete"`
	IdRole    string `json:"id_role"`
}

func (v *RolePermissions) BeforeCreate(tx *gorm.DB) error {
	v.Id = helpers.CreateId()
	return nil
}
