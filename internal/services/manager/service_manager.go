package manager

import (
	"goproject/internal/database/db"
	postController "goproject/internal/modules/post/controller"
	postRepo "goproject/internal/modules/post/repository"
	postService "goproject/internal/modules/post/service"
	userController "goproject/internal/modules/user/controller"
	userRepo "goproject/internal/modules/user/repository"
	userService "goproject/internal/modules/user/service"
)

type ServiceManager struct {
	UserService    *userService.UserService
	PostService    *postService.PostService
	UserController *userController.UserController
	PostController *postController.PostController
}

func NewServiceManager(conn *db.MongoDBConnection) *ServiceManager {
	userRepository := userRepo.NewUserRepository(conn.DB)
	postRepository := postRepo.NewPostRepository(conn.DB)

	uService := userService.NewUserService(userRepository)
	pService := postService.NewPostService(postRepository, uService, conn.DB)

	uc := userController.NewUserController(uService)
	pc := postController.NewPostController(pService, uService, conn.DB)

	return &ServiceManager{
		UserService:    uService,
		PostService:    pService,
		UserController: uc,
		PostController: pc,
	}
}
