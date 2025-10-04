package auth

import (
	"time"
)

// ------------------- Login -------------------
type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ------------------- Register -------------------
type RegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

// ------------------- User DTO -------------------
type UserDTO struct {
	ID            uint      `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Role          string    `json:"role"`   //
	Status        string    `json:"status"` // active/inactive
	CreatedAt     time.Time `json:"created_at"`
}

// ------------------- Login Response -------------------
type LoginResponse struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *UserDTO
}

// ------------------- Password Change -------------------
type PasswordChange struct {
	Password        string `form:"password" json:"password" binding:"required"`                                     // current password
	NewPassword     string `form:"new_password" json:"new_password" binding:"required"`                             // new password
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=NewPassword"` // confirm password must match
}
