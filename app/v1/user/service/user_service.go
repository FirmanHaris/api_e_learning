// nama package sesuai nama folder
package service

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/repository"
	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/utils/r"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// service berisi logika sebelum disimpan ke database
type UserService interface {
	GetAllUser(ctx context.Context) ([]*domain.User, r.Ex)
	GetUserById(ctx context.Context, id primitive.ObjectID) (*domain.User, r.Ex)
	AddUser(ctx context.Context, data *domain.User) (primitive.ObjectID, r.Ex)
	UpdateUser(ctx context.Context, data *domain.User) r.Ex
	DeleteUser(ctx context.Context, data *domain.User) r.Ex
}

type baseUserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &baseUserService{userRepo: userRepo}
}

func (b *baseUserService) GetAllUser(ctx context.Context) ([]*domain.User, r.Ex) {
	return b.userRepo.GetAll(ctx)
}

func (b *baseUserService) GetUserById(ctx context.Context, id primitive.ObjectID) (*domain.User, r.Ex) {
	return b.userRepo.GetById(ctx, id)
}

func (b *baseUserService) AddUser(ctx context.Context, data *domain.User) (primitive.ObjectID, r.Ex) {
	return b.userRepo.Add(ctx, data)
}
func (b *baseUserService) UpdateUser(ctx context.Context, data *domain.User) r.Ex {
	return b.userRepo.Update(ctx, data)
}
func (b *baseUserService) DeleteUser(ctx context.Context, data *domain.User) r.Ex {
	return b.userRepo.Delete(ctx, data)
}
