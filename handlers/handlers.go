package handlers

import (
	"go-mysql-gorm-gin/database"
	"go-mysql-gorm-gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	var users []model.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	var user model.User
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid i"})
		return
	}
	if err := database.DB.First(&user, (uint)(id)).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user model.User

	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to fetch User"})
		return
	}
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Name = input.Name
	user.Age = input.Age
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	if err := database.DB.Delete(&model.User{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "User Deleted successfully!")
}
