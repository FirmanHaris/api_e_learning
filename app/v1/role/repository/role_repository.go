// nama package sesuai dengan nama folder
package repository

import (
	"context"

	"github.com/FirmanHaris/api_e_learning/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository berisi crud ke database
type RoleRepository interface {
	// untuk bisa dibaca di service fungsi yg sudah dibuat harus di definikan di interface
	GetAll(ctx context.Context) ([]*domain.Role, error)
}

type baseRoleRepository struct {
	coll *mongo.Collection
}

func NewRoleRepository(coll *mongo.Database) RoleRepository {
	role := coll.Collection("role")
	return &baseRoleRepository{coll: role}
}

// fungsi untuk mengconvert monggo cursor ke struct Role digunakan jika mereturn data array dan tidak perlu di definikan di interface
func curration(ctx context.Context, cursor *mongo.Cursor) ([]*domain.Role, error) {
	var result []*domain.Role
	for cursor.Next(ctx) {
		var role domain.Role
		err := cursor.Decode(&role)
		if err != nil {
			return nil, err
		}
		result = append(result, &role)
	}
	return result, nil
}

// fungsi untuk mengambil semua data role dari database
func (b *baseRoleRepository) GetAll(ctx context.Context) ([]*domain.Role, error) {
	filter := bson.M{}
	role, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return curration(ctx, role)
}
