package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	RoleID   primitive.ObjectID `json:"role_id" bson:"role_id"`
	Log      Log                `json:"log" bson:"log"`
}
