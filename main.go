package main

import (
	"go-mysql-gorm-gin/database"
	"go-mysql-gorm-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	users := r.Group("/api/v1/users")
	{
		users.GET("/", handlers.GetAllUsers)
		users.GET("/:id", handlers.GetUserById)
		users.POST("/", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
	r.Run(":8080")
}
