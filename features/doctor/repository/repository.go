package repository

import (
	"errors"
	"qhealth/domain"
	"qhealth/features/doctor"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) doctor.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateDoctor(doctor domain.Doctor) error {
	err := r.db.Create(&doctor).Error
	if err != nil {
		return nil
	}
	return nil
}

func (r *repository) FindByEmail(email string) (domain.Doctor, error) {
	doctor := domain.Doctor{}
	err := r.db.Where("email = ?", email).First(&doctor).Error
	if err != nil {
		return domain.Doctor{}, err
	}
	return doctor, nil
}

func (r *repository) FindById(id string) (domain.Doctor, error) {
	doctor := domain.Doctor{}
	err := r.db.Where("id = ?", id).First(&doctor).Error
	if err != nil {
		return domain.Doctor{}, err
	}
	return doctor, nil
}

func (r *repository) UpdatePass(email, newPass string) error {
	doctor := domain.Doctor{}

	if err := r.db.Where("email = ?", email).First(&doctor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("doctor not found")
		}
		return err
	}

	doctor.Password = newPass
	if err := r.db.Save(&doctor).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteProfile(email string) error {
	err := r.db.Where("email = ?", email).Delete(&domain.Doctor{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateProfile(email string, doctor domain.Doctor) error {
	_, err := r.FindByEmail(email)
	if err != nil {
		return err
	}

	err = r.db.Where("email = ?", email).Updates(&doctor).Error
	if err != nil {
		return err
	}
	
	return nil
}