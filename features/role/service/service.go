package service

import (
	"qhealth/domain"
	"qhealth/features/role"
)

type service struct {
	repo role.Repository
}

func NewRoleService(repo role.Repository) role.Service {
	return &service{repo: repo}
}

func (s *service) CreateRole(roleReq domain.RoleReq) error {
	role := domain.ReqToRole(roleReq)

	err := s.repo.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllRole() ([]domain.RoleResp, error) {
	role, err := s.repo.GetAllRole()

	if err != nil {
		return nil, err
	}

	result := domain.ListRoleToResp(role)
	return result, nil
}