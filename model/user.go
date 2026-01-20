package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0"`
}
