package router

import (
	"fmt"
	"os"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/controller"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/middleware"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	// Initiate router v1
	router := gin.Default()
	v1 := router.Group("/v1")

	// Create cookie
	store := cookie.NewStore([]byte("secret"))
	v1.Use(sessions.Sessions("mysession", store))

	// Initiate user model
	userModel := model.NewUser(db)
	photoModel := model.NewPhoto(db)

	// Initiate user service
	userService := service.NewUser(userModel)
	photoService := service.NewPhoto(photoModel)

	// Initiate controller
	auth := controller.NewAuth(userService)
	user := controller.NewUser(userService)
	photo := controller.NewPhoto(photoService)

	// User Router
	userRouter := v1.Group("/users")

	userRouter.POST("/register", auth.Register)
	userRouter.POST("/login", auth.Login)
	userRouter.GET("/logout", auth.LogOut)

	// Set middleware is login true
	userRouter.Use(middleware.ChekLogin())
	userRouter.GET("/generate-code", auth.GenerateCode)
	userRouter.POST("/validate-code", auth.ValidateCode)
	userRouter.PUT("/:id", user.UpdateUser)
	userRouter.DELETE("/:id", user.DeleteUser)

	photosRouter := v1.Group("/photos")
	photosRouter.Use(middleware.ChekLogin())

	photosRouter.POST("/", photo.CreatePhoto)
	photosRouter.GET("/", photo.GetPhoto)
	photosRouter.PUT("/:id", photo.UpdatePhoto)
	photosRouter.DELETE("/:id", photo.DeletePhoto)

	router.Run(os.Getenv("DB_HOST") + ":8080")
	fmt.Println("Running on " + os.Getenv("DB_HOST") + " : " + os.Getenv("DB_PORT"))
}
