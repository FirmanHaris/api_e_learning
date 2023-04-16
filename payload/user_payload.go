package payload

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterUserPayload struct {
	Username string             `json:"username" form:"username" validate:"required"`
	Email    string             `json:"email"  form:"email" validate:"required,email"`
	Password string             `json:"password" form:"password" validate:"required,gte=8"`
	RoleID   primitive.ObjectID `json:"role_id" form:"role_id" validate:"required"`
}

type UserGetByID struct {
	ID primitive.ObjectID `json:"id" query:"id" param:"id" form:"id" validate:"required"`
}

type UpdatePasswordPayload struct {
	UserGetByID
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
	OldPassowrd string `json:"old_passowrd" form:"old_password" validate:"required"`
}
