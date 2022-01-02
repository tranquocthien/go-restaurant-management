package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	Name       *string            `json:"name" bson:"name" validate:"required, min=2, max=100"`
	Price      *float64           `json:"price" bson:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Deleted_at time.Time          `json:"deleted_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}
