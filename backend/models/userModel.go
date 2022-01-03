package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"` // tag golang
	First_name    *string            `json:"first_name" validate:"required"`
	Last_name     *string            `json:"last_name" validate:"required"`
	Password      *string            `json:"Password" validate:"required"`
	Email         *string            `json:"email" validate:"required"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Deleted_at    time.Time          `json:"deleted_at"`
}
