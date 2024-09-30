package service

import (
	"errors"
	"qhealth/domain"
	"qhealth/features/users"
	"qhealth/helpers"
	"qhealth/helpers/middleware"

	"github.com/google/uuid"
)

type service struct {
	repo users.Repository
}

func NewService(repo users.Repository) users.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(userReq domain.UserRegister) error {
    role, err := s.repo.GetRoleByName("user")
    if err != nil {
        return err
    }

    user := domain.UserRegisterToUser(userReq)
    user.IdRole = role.Id 

    err = s.repo.CreateUser(user)
    if err != nil {
        return err
    }

    return nil
}

// func (s *service) RegisterUser(user *domain.UserRegister) error {
// 	role, err := s.repo.GetRoleByName("user")
// 	if err != nil {
// 		return err
// 	}
// 	user.IdRole = role.Id

// 	return s.repo.CreateUser(domain.User{})
// }

func (s *service) Login(userReq domain.UserLogin) (string, error) {
	user, err := s.repo.FindByEmail(userReq.Email)
	if err != nil {
		return "", err
	}

	ok, err := helpers.ComparePass([]byte(user.Password), []byte(userReq.Password))
	if !ok {
		if err != nil {
			return "", errors.New("invalid password")
		}
	}

	token, err := middleware.CreateToken(user.Id, user.Email) 
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) ChangePass(email string, reqPass domain.ReqChangePass) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	ok, err := helpers.ComparePass([]byte(user.Password), []byte(reqPass.OldPass))
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
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	if user.Email != email {
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

func (s *service) DeleteUser(email string) error {
	return s.repo.DeleteUser(email)
}

func (s *service) UpdateUser(email string, user domain.UserReq) error {
	data := domain.ReqToUser(user)
	err := s.repo.UpdateUser(email, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) InitializeRolesAndPermission() error {
    roles := []domain.Role{
        {Id: uuid.New().String(), Name: "admin"},
        {Id: uuid.New().String(), Name: "user"},
    }

    for i, role := range roles {
        if err := s.repo.CreateRole(&role); err != nil {
            return err
        }
        roles[i] = role
    }

    permission := []domain.RolePermissionResp{
        {Id: uuid.New().String(), CanCreate: true, CanRead: true, CanEdit: true, CanDelete: true, IdRole: roles[0].Id},  
        {Id: uuid.New().String(), CanCreate: false, CanRead: true, CanEdit: false, CanDelete: false, IdRole: roles[1].Id}, 
    }

    for _, perm := range permission {
        rolePerm := domain.RolePermissionRespToRolePermission(perm)
        if err := s.repo.CreateRolePermission(&rolePerm); err != nil {
            return err
        }
    }

    return nil
}
