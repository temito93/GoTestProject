package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject/internal/modules/post/dto"
	"goproject/internal/modules/post/service"
	userService "goproject/internal/modules/user/service"
	"net/http"
)

type PostController struct {
	postService *service.PostService
	userService *userService.UserService
	db          *mongo.Database
}

func NewPostController(sm *service.PostService, userService *userService.UserService, db *mongo.Database) *PostController {
	return &PostController{postService: sm, userService: userService, db: db}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req dto.CreatePostDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := pc.postService.CreatePost(c.Request.Context(), req.Title, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"post": post})
}
