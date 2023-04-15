// nama package sesuai nama folder
package service

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/repository"
	"github.com/FirmanHaris/api_e_learning/utils/r"
)

// service berisi logika sebelum disimpan ke database
type RoleService interface {
	GetAllRole(ctx context.Context) ([]*domain.Role, r.Ex)
}

type baseRoleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &baseRoleService{roleRepo: roleRepo}
}

func (b *baseRoleService) GetAllRole(ctx context.Context) ([]*domain.Role, r.Ex) {
	return b.roleRepo.GetAll(ctx)
}
