// nama package sesuai nama folder
package service

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/role/repository"
	"github.com/FirmanHaris/api_e_learning/domain"
)

// service berisi logika sebelum disimpan ke database
type RoleService interface {
	GetAllRole(ctx context.Context) ([]*domain.Role, error)
}

type baseRoleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &baseRoleService{roleRepo: roleRepo}
}

func (b *baseRoleService) GetAllRole(ctx context.Context) ([]*domain.Role, error) {
	return b.roleRepo.GetAll(ctx)
}