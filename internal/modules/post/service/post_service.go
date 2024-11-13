package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject/internal/modules/post/entity"
	"goproject/internal/modules/post/repository"
	userService "goproject/internal/modules/user/service"
)

type PostService struct {
	postRepo    *repository.PostRepository
	userService *userService.UserService
	db          *mongo.Database
}

func NewPostService(postRepo *repository.PostRepository, userService *userService.UserService, db *mongo.Database) *PostService {
	return &PostService{
		postRepo:    postRepo,
		userService: userService,
		db:          db,
	}
}

func (s *PostService) CreatePost(ctx context.Context, title, userID string) (*entity.PostEntity, error) {
	//s.userService.CreateUser()
	post := &entity.PostEntity{Title: title, UserID: userID}
	err := s.postRepo.CreatePost(ctx, post)
	return post, err
}
