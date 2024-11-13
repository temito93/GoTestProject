package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject/internal/modules/user/entity"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	result, err := r.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	// Set the ID of the user to the inserted ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	} else {
		return nil, errors.New("cannot convert to object id")
	}

	return user, err
}

func Testnet() {}
