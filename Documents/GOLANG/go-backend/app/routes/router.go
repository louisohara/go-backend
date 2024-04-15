package routes 

import (
	"database/sql"

	"example.com/app/controllers"
	"example.com/app/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo) 

	router.GET("/users", userController.getUsers)
	router.GET("/users/:id", userController.getUserByID)
	router.PUT("/users/:id", userController.updateUser)
	router.POST("/users", userController.postUsers)
	router.DELETE("/users/:id", userController.deleteUser)

	return router
	// router.Run("localhost:8080")
}