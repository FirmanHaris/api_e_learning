package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
