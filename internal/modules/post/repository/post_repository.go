package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject/internal/modules/post/entity"
)

type PostRepository struct {
	db *mongo.Database
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *entity.PostEntity) error {
	_, err := r.db.Collection("posts").InsertOne(ctx, post)
	return err
}
