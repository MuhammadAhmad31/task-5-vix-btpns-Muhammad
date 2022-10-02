package main

import (
	"WebApiGo/controllers"
	"WebApiGo/database"
	"WebApiGo/helpers"
	"WebApiGo/middleware"
	"WebApiGo/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = database.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	photoRepository repository.PhotoRepository = repository.NewPhotoRepository(db)
	jwtService      helpers.JWTHelpers         = helpers.NewJWTHelpers()
	userService     helpers.UserHelpers        = helpers.NewUserHelpers(userRepository)
	photoService    helpers.PhotoHelpers       = helpers.NewPhotoHelpers(photoRepository)
	authService     helpers.AuthHelpers        = helpers.NewAuthHelpers(userRepository)
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	photoController controller.PhotoController = controller.NewPhotoController(photoService, jwtService)
)

func main() {
	defer database.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/users/login", userController.Profile)
		userRoutes.PUT("/users/:userID", userController.Update)
	}

	photoRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/photos", photoController.All)
		photoRoutes.POST("/photos", photoController.Insert)
		photoRoutes.PUT("/:photoID", photoController.Update)
		photoRoutes.DELETE("/:photoID", photoController.Delete)
	}
	r.Run()
}
