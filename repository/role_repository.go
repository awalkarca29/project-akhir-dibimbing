package repository

import (
	"project-akhir-awal/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(role entity.Role) (entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Save(role entity.Role) (entity.Role, error) {
	err := r.db.Create(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}
