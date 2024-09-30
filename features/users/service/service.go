package service

import (
	"errors"
	"qhealth/domain"
	"qhealth/features/users"
	"qhealth/helpers"
	"qhealth/helpers/middleware"
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

func (s *service) RegisterAdmin(adminReq domain.UserRegister) error {
    role, err := s.repo.GetRoleByName("admin")
    if err != nil {
        return err
    }

    admin := domain.UserRegisterToUser(adminReq)
    admin.IdRole = role.Id 

    err = s.repo.CreateUser(admin)
    if err != nil {
        return err
    }

    return nil
}

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
        {Name: "admin"},
        {Name: "user"},
    }

    for _, role := range roles {
        var existingRole domain.Role
        if err := s.repo.FindRoleByName(role.Name, &existingRole); err != nil {
            if err := s.repo.CreateRole(&role); err != nil {
                return err
            }
        }
    }

    var adminRole, userRole domain.Role
    if err := s.repo.FindRoleByName("admin", &adminRole); err != nil {
        return err
    }
    if err := s.repo.FindRoleByName("user", &userRole); err != nil {
        return err
    }

    adminPermissions := domain.RolePermissions{
        CanCreate: true,
        CanRead:   true,
        CanEdit:   true,
        CanDelete: true,
        IdRole:    adminRole.Id,
    }

    userPermissions := domain.RolePermissions{
        CanCreate: false,
        CanRead:   true,
        CanEdit:   true, 
        CanDelete: true, 
        IdRole:    userRole.Id,
    }

    var existingAdminPerm domain.RolePermissions
    if err := s.repo.FindRolePermissionByRoleId(adminRole.Id, &existingAdminPerm); err != nil {
        if err := s.repo.CreateRolePermission(&adminPermissions); err != nil {
            return err
        }
    } else {
        existingAdminPerm.CanCreate = adminPermissions.CanCreate
        existingAdminPerm.CanRead = adminPermissions.CanRead
        existingAdminPerm.CanEdit = adminPermissions.CanEdit
        existingAdminPerm.CanDelete = adminPermissions.CanDelete

        if err := s.repo.UpdateRolePermission(&existingAdminPerm); err != nil {
            return err
        }
    }

    var existingUserPerm domain.RolePermissions
    if err := s.repo.FindRolePermissionByRoleId(userRole.Id, &existingUserPerm); err != nil {
        if err := s.repo.CreateRolePermission(&userPermissions); err != nil {
            return err
        }
    } else {
        existingUserPerm.CanCreate = userPermissions.CanCreate
        existingUserPerm.CanRead = userPermissions.CanRead
        existingUserPerm.CanEdit = userPermissions.CanEdit
        existingUserPerm.CanDelete = userPermissions.CanDelete

        if err := s.repo.UpdateRolePermission(&existingUserPerm); err != nil {
            return err
        }
    }

    return nil
}