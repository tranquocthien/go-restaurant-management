package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	Name       string
	Category   string
	Start_Date *time.Time
	End_Date   *time.Time
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
	Menu_id    string
}
