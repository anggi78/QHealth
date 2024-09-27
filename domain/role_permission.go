package domain

type RolePermission struct {
    Id        string `gorm:"primaryKey"`
    CanCreate bool   `gorm:"default:false"`
    CanRead   bool   `gorm:"default:false"`
    CanEdit   bool   `gorm:"default:false"`
    CanDelete bool   `gorm:"default:false"`
    IdRole    string 
	Role      Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

