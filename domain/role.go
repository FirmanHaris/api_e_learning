package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Role string             `json:"role" bson:"role"`
	Log  Log                `json:"log" bson:"log"`
}
