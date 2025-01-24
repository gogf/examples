package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	CreatedAt int64              `bson:"created_at,omitempty"`
	UpdatedAt int64              `bson:"updated_at,omitempty"`
}
