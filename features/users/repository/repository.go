package repository

import (
	"errors"
	"qhealth/domain"
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
	err := r.db.Preload("Role").Create(&user).Error
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

func (r *repository) FindById(id string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// func (r *repository) FindCodeByEmail(email string) (string, error) {
// 	user := domain.User{}
// 	err := r.db.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		return "", errors.New("not found")
// 	}
// 	return user.Code, nil
// }

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

func (r *repository) DeleteUser(email string) error {
	err := r.db.Where("email = ?", email).Delete(&domain.User{}).Error
	if err != nil {
		return err
	}

	// Contoh jika ada entitas related:
	// err = r.db.Where("user_id = ?", userID).Delete(&domain.RelatedEntity{}).Error
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r *repository) UpdateUser(email string, user domain.User) error {
	_, err := r.FindByEmail(email)
	if err != nil {
		return err
	}

	err = r.db.Where("email = ?", email).Updates(&user).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (r *repository) GetRoleByName(name string) (domain.Role, error) {
    var role domain.Role
    err := r.db.Where("name = ?", name).First(&role).Error
    if err != nil {
        return role, err
    }
    return role, nil
}

func (r *repository) FindRoleByName(name string, role *domain.Role) error {
    return r.db.Where("name = ?", name).First(role).Error
}

func (r *repository) FindRolePermissionByRoleId(roleId string, permission *domain.RolePermissions) error {
    return r.db.Where("id_role = ?", roleId).First(permission).Error
}

func (r *repository) CreateRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *repository) CreateRolePermission(permission *domain.RolePermissions) error {
    return r.db.Create(permission).Error
}

func (r *repository) UpdateRolePermission(permission *domain.RolePermissions) error {
    return r.db.Save(permission).Error
}