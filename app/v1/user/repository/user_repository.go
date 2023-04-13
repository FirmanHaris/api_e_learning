// nama package sesuai dengan nama folder
package repository

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository berisi crud ke database
type UserRepository interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
}

type baseUserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) UserRepository {
	return &baseUserRepository{coll: coll}
}

func curration(ctx context.Context, cursor *mongo.Cursor) ([]*domain.User, error) {
	var result []*domain.User
	for cursor.Next(ctx) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, &user)
	}
	return result, nil
}

func (b *baseUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	filter := bson.M{}
	user, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return curration(ctx, user)
}
