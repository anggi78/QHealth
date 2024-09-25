package repository

import (
	"qhealth/domain"
	"qhealth/features/role"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateRole(role domain.Role) error {
	err := r.db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllRole() ([]domain.Role, error) {
	var role []domain.Role
	err := r.db.Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}