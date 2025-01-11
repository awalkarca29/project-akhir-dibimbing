package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"
)

type CreateRoleInput struct {
	Name string `json:"name" binding:"required"`
}

type RoleService interface {
	CreateRole(input CreateRoleInput) (entity.Role, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *roleService {
	return &roleService{roleRepository}
}

func (s *roleService) CreateRole(input CreateRoleInput) (entity.Role, error) {
	role := entity.Role{}
	role.Name = input.Name

	newRole, err := s.roleRepository.Save(role)
	if err != nil {
		return newRole, err
	}

	return newRole, nil
}
