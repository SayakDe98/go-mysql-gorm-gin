package main

import (
	"go-mysql-gorm-gin/database"
	handlers "go-mysql-gorm-gin/handler"
	"go-mysql-gorm-gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	// users := r.Group("/api/v1/users")
	// {
	// 	users.GET("/", handlers.GetAllUsers)
	// 	users.GET("/:id", handlers.GetUserById)
	// 	users.POST("/", handlers.CreateUser)
	// 	users.PUT("/:id", handlers.UpdateUser)
	// 	users.DELETE("/:id", handlers.DeleteUser)
	// }
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.LoginHandler)
	}

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/:id", handlers.GetUserById)
		userRoutes.GET("/", handlers.GetAllUsers)
	}
	r.Run(":8080")
}
