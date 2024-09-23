package repository

import (
	"errors"
	"qhealth/entity/domain"
	"qhealth/features/users"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user domain.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil
	}
	return nil
}

func (r *repository) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) FindCodeByEmail(email string) (string, error) {
	user := domain.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", errors.New("not found")
	}
	return user.Code, nil
}

func (r *repository) UpdatePass(email, newPass string) error {
	user := domain.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	user.Password = newPass
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}