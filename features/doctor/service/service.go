package service

import (
	"errors"
	"qhealth/domain"
	"qhealth/features/doctor"
	"qhealth/helpers"
	"qhealth/helpers/middleware"
)

type service struct {
	repo doctor.Repository
}

func NewDoctorService(repo doctor.Repository) doctor.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(doctorReq domain.DoctorRegister) error {
	doctor := domain.DoctorRegisterToDoctor(doctorReq)
	err := s.repo.CreateDoctor(doctor)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Login(doctorReq domain.DoctorLogin) (string, error) {
	doctor, err := s.repo.FindByEmail(doctorReq.Email)
	if err != nil {
		return "", err
	}

	ok, err := helpers.ComparePass([]byte(doctor.Password), []byte(doctorReq.Password))
	if !ok {
		if err != nil {
			return "", errors.New("invalid password")
		}
	}

	token, err := middleware.CreateToken(doctor.Id, doctor.Email) 
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) ChangePass(email string, reqPass domain.ReqChangePassDoctor) error {
	doctor, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	ok, err := helpers.ComparePass([]byte(doctor.Password), []byte(reqPass.OldPass))
	if !ok {
		return errors.New("old password doesn't match")
	}
	if reqPass.ConfirmPass != reqPass.ConfirmPass {
		return errors.New("confirm password is doesn't match")
	}
	if err != nil {
		return err
	}

	hassPass, err := helpers.HassPass(reqPass.NewPass) 
	if err != nil {
		return err
	}

	err = s.repo.UpdatePass(email, hassPass)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChangePassForgot(email, newPass string) error {
	hassPass, err := helpers.HassPass(newPass)
	if err != nil {
		return err
	}

	err = s.repo.UpdatePass(email, hassPass)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ForgotPassword(email string) error {
	doctor, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	if doctor.Email != email {
		return errors.New("invalid email")
	}

	// code, err := s.repo.FindCodeByEmail(email)
	// if err != nil {
	// 	return err
	// }

	// err = helpers.SendEmail(email, "code verification", code)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *service) DeleteProfile(email string) error {
	return s.repo.DeleteProfile(email)
}

func (s *service) UpdateProfile(email string, doctor domain.DoctorReq) error {
	data := domain.ReqToDoctor(doctor)
	err := s.repo.UpdateProfile(email, data)
	if err != nil {
		return err
	}
	return nil
}