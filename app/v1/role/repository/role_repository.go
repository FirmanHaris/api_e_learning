// nama package sesuai dengan nama folder
package repository

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/utils/r"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository berisi crud ke database
type RoleRepository interface {
	// untuk bisa dibaca di service fungsi yg sudah dibuat harus di definikan di interface
	GetAll(ctx context.Context) ([]*domain.Role, r.Ex)
}

type baseRoleRepository struct {
	coll *mongo.Collection
}

func NewRoleRepository(coll *mongo.Database) RoleRepository {
	role := coll.Collection("role")
	return &baseRoleRepository{coll: role}
}

// fungsi untuk mengconvert monggo cursor ke struct Role digunakan jika mereturn data array dan tidak perlu di definikan di interface
func (b *baseRoleRepository) consumeCursor(ctx context.Context, cursor *mongo.Cursor) ([]*domain.Role, r.Ex) {
	var result []*domain.Role
	for cursor.Next(ctx) {
		var role domain.Role
		err := cursor.Decode(&role)
		if err != nil {
			return nil, r.NewErrorMongo(b.coll.Name(), err)
		}
		result = append(result, &role)
	}
	return result, nil
}

// fungsi untuk mengambil semua data role dari database
func (b *baseRoleRepository) GetAll(ctx context.Context) ([]*domain.Role, r.Ex) {
	filter := bson.M{}
	role, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, r.NewErrorMongo(b.coll.Name(), err)
	}
	if role == nil {
		return nil, r.NewErrorDataNotFound(b.coll.Name(), err)
	}
	result, error := b.consumeCursor(ctx, role)
	return result, error
}
