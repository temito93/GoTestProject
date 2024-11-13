package index

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateUserIndex(db *mongo.Database) error {
	collection := db.Collection("users")
	trueVal := true
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: &options.IndexOptions{
				Unique: &trueVal,
			},
		},
	)
	if err != nil {
		log.Fatalf("Error creating unique index on email: %v", err)
		return err
	}
	log.Println("Unique index created on 'email' field.")
	return nil
}
