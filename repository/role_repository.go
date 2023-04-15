// nama package sesuai dengan nama folder
package repository

import (
	"context"
	"time"

	"github.com/FirmanHaris/api_e_learning/domain"
	"github.com/FirmanHaris/api_e_learning/utils/r"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// repository berisi crud ke database
type RoleRepository interface {
	// untuk bisa dibaca di service fungsi yg sudah dibuat harus di definikan di interface
	GetAll(ctx context.Context) ([]*domain.Role, r.Ex)
	Add(ctx context.Context, d *domain.Role) (primitive.ObjectID, r.Ex)
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.Role, r.Ex)
	Update(ctx context.Context, d *domain.Role) r.Ex
	Delete(ctx context.Context, d *domain.Role) r.Ex
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

func (p *baseRoleRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (p *baseRoleRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.Role, r.Ex) {
	res := new(domain.Role)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *baseRoleRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}

	return nil
}

// fungsi untuk mengambil semua data dari database
func (b *baseRoleRepository) GetAll(ctx context.Context) ([]*domain.Role, r.Ex) {
	filter := bson.M{}
	role, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, r.NewErrorMongo(b.coll.Name(), err)
	}
	return b.consumeCursor(ctx, role)
}

func (p *baseRoleRepository) Add(ctx context.Context, d *domain.Role) (primitive.ObjectID, r.Ex) {
	now := time.Now().UTC()
	dateNow := primitive.NewDateTimeFromTime(now)
	d.Log = domain.Log{
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}
	return p.insert(ctx, d)
}

func (p *baseRoleRepository) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Role, r.Ex) {
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	return p.findOne(ctx, filter)
}

func (p *baseRoleRepository) Update(ctx context.Context, d *domain.Role) r.Ex {
	now := time.Now().UTC()
	filter := bson.D{
		{Key: "_id", Value: d.ID},
	}
	d.Log.UpdatedAt = primitive.NewDateTimeFromTime(now)
	updated := bson.D{
		{Key: "$set", Value: d},
	}
	return p.updatedOne(ctx, filter, updated)
}

func (p *baseRoleRepository) Delete(ctx context.Context, d *domain.Role) r.Ex {
	filter := bson.D{
		{Key: "_id", Value: d.ID},
	}
	_, err := p.coll.DeleteOne(ctx, filter)
	if err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}
