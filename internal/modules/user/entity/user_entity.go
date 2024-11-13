package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
