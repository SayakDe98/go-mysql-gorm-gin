package handlers

import (
	"go-mysql-gorm-gin/auth"
	"go-mysql-gorm-gin/database"
	"go-mysql-gorm-gin/dto"
	"go-mysql-gorm-gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var dto dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, err := auth.HashPassword(dto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	user := model.User{
		Name:     dto.Name,
		Age:      dto.Age,
		Email:    dto.Email,
		Password: hashed,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}
