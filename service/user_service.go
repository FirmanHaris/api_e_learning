// nama package sesuai nama folder
package service

import (
	"context"
	"time"

	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/payload"
	"github.com/FirmanHaris/api_e_learning/repository"
	"github.com/FirmanHaris/api_e_learning/utils/password"
	"github.com/FirmanHaris/api_e_learning/utils/r"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// service berisi logika sebelum disimpan ke database
type UserService interface {
	GetAllUser(ctx context.Context) ([]*domain.User, r.Ex)
	GetUserById(ctx context.Context, id primitive.ObjectID) (*domain.User, r.Ex)
	AddUser(ctx context.Context, payload *payload.RegisterUserPayload) (primitive.ObjectID, r.Ex)
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

func (b *baseUserService) AddUser(ctx context.Context, payload *payload.RegisterUserPayload) (primitive.ObjectID, r.Ex) {
	hashPassword, err := password.HashPassword(payload.Password)
	if err != nil {
		return primitive.NilObjectID, r.NewErr(err.Error())
	}
	timeNow := primitive.NewDateTimeFromTime(time.Now().UTC())
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: payload.Username,
		Password: hashPassword,
		Email:    payload.Email,
		RoleID:   payload.RoleID,
		Log: domain.Log{
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
	}
	return b.userRepo.Add(ctx, user)
}
func (b *baseUserService) UpdateUser(ctx context.Context, data *domain.User) r.Ex {
	return b.userRepo.Update(ctx, data)
}
func (b *baseUserService) DeleteUser(ctx context.Context, data *domain.User) r.Ex {
	return b.userRepo.Delete(ctx, data)
}
