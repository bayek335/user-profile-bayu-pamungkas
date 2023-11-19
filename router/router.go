package router

import (
	"fmt"
	"os"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/controller"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	// Initiate router v1
	router := gin.Default()
	v1 := router.Group("/v1")

	// Initiate user model
	userModel := model.NewUser(db)

	// Initiate user model
	userService := service.NewUser(userModel)

	// Initiate user model
	auth := controller.NewAuth(userService)

	// User Router
	userRouter := v1.Group("/users")

	userRouter.POST("/register", auth.Register)

	router.Run(os.Getenv("DB_HOST") + ":8080")
	fmt.Println("Running on " + os.Getenv("DB_HOST") + " : " + os.Getenv("DB_PORT"))
}
