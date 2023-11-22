package app

import "time"

// Struct that implement table in database
type User struct {
	ID        int       `db:"id" gorm:"primaryKey"`
	Username  string    `db:"username"`
	Email     string    `db:"email" gorm:"unique"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Struct that bind the user request while register
type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	IsActive bool   `db:"is_active"`
}

// Struct that bind the user request  while login
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// Struct that formating the json response success
type UserResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"data"`
}

// Struct that formating the json response fail
type UserResponseFail struct {
	Success bool   `json:"succes"`
	Message string `json:"message"`
}
