package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex"`
	Password string `json:"-"`
}
