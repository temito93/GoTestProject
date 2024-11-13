package routes

import (
	"github.com/gin-gonic/gin"
	"goproject/internal/database/db"
	"goproject/internal/services/manager"
)

func SetupRoutes(router *gin.Engine, conn *db.MongoDBConnection) {
	serviceManager := manager.NewServiceManager(conn)

	router.POST("/users", serviceManager.UserController.CreateUser)
	router.POST("/posts", serviceManager.PostController.CreatePost)
}
