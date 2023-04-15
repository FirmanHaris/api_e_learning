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
type UserRepository interface {
	// untuk bisa dibaca di service fungsi yg sudah dibuat harus di definikan di interface
	Add(ctx context.Context, d *domain.User) (primitive.ObjectID, r.Ex)
	GetAll(ctx context.Context) ([]*domain.User, r.Ex)
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.User, r.Ex)
	Update(ctx context.Context, d *domain.User) r.Ex
	Delete(ctx context.Context, d *domain.User) r.Ex
}

type baseUserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Database) UserRepository {
	user := coll.Collection("user")
	return &baseUserRepository{coll: user}
}

// fungsi untuk mengconvert monggo cursor ke struct User digunakan jika mereturn data array dan tidak perlu di definikan di interface
func (b *baseUserRepository) consumeCursor(ctx context.Context, cursor *mongo.Cursor) ([]*domain.User, r.Ex) {
	var result []*domain.User
	for cursor.Next(ctx) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, r.NewErrorMongo(b.coll.Name(), err)
		}
		result = append(result, &user)
	}
	return result, nil
}

func (p *baseUserRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (p *baseUserRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.User, r.Ex) {
	res := new(domain.User)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *baseUserRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}

	return nil
}

// fungsi untuk mengambil semua data dari database
func (b *baseUserRepository) GetAll(ctx context.Context) ([]*domain.User, r.Ex) {
	filter := bson.M{}
	user, err := b.coll.Find(ctx, filter)
	if err != nil {
		return nil, r.NewErrorMongo(b.coll.Name(), err)
	}
	return b.consumeCursor(ctx, user)
}

func (p *baseUserRepository) Add(ctx context.Context, d *domain.User) (primitive.ObjectID, r.Ex) {
	now := time.Now().UTC()
	dateNow := primitive.NewDateTimeFromTime(now)
	d.Log = domain.Log{
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
	}
	return p.insert(ctx, d)
}

func (p *baseUserRepository) GetById(ctx context.Context, id primitive.ObjectID) (*domain.User, r.Ex) {
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	return p.findOne(ctx, filter)
}

func (p *baseUserRepository) Update(ctx context.Context, d *domain.User) r.Ex {
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

func (p *baseUserRepository) Delete(ctx context.Context, d *domain.User) r.Ex {
	filter := bson.D{
		{Key: "_id", Value: d.ID},
	}
	_, err := p.coll.DeleteOne(ctx, filter)
	if err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}
