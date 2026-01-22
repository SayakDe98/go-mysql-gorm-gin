package dto

type RegisterUserDTO struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=2"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
