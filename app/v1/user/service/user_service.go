// nama package sesuai nama folder
package service

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/app/v1/user/repository"
	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/utils/r"
)

// service berisi logika sebelum disimpan ke database
type UserService interface {
	GetAllUser(ctx context.Context) ([]*domain.User, r.Ex)
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
