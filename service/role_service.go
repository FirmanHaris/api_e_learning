// nama package sesuai nama folder
package service

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/repository"
	"github.com/FirmanHaris/api_e_learning/utils/r"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// service berisi logika sebelum disimpan ke database
type RoleService interface {
	GetAllRole(ctx context.Context) ([]*domain.Role, r.Ex)
	GetRoleById(ctx context.Context, id primitive.ObjectID) (*domain.Role, r.Ex)
	AddRole(ctx context.Context, data *domain.Role) (primitive.ObjectID, r.Ex)
	UpdateRole(ctx context.Context, data *domain.Role) r.Ex
	DeleteRole(ctx context.Context, data *domain.Role) r.Ex
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

func (b *baseRoleService) GetRoleById(ctx context.Context, id primitive.ObjectID) (*domain.Role, r.Ex) {
	return b.roleRepo.GetById(ctx, id)
}

func (b *baseRoleService) AddRole(ctx context.Context, data *domain.Role) (primitive.ObjectID, r.Ex) {
	return b.roleRepo.Add(ctx, data)
}
func (b *baseRoleService) UpdateRole(ctx context.Context, data *domain.Role) r.Ex {
	return b.roleRepo.Update(ctx, data)
}
func (b *baseRoleService) DeleteRole(ctx context.Context, data *domain.Role) r.Ex {
	return b.roleRepo.Delete(ctx, data)
}
