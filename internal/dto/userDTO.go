package dto

import (
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/models"
)

// Used for user login request
type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Used for user registration
type RegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

type UserDTO struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Role          string `json:"role"`   // usually "user"
	Status        string `json:"status"` // active/inactive
}

type LoginResponse struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *models.User // change into userDTO
}

type ProfileDTO struct {
	ID      uint       `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Role    string     `json:"role"`
	Address AddressDTO `json:"address"`
}

type AddressDTO struct {
	ID          uint   `json:"id"`
	AddressLine string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"zip_code"`
	Country     string `json:"country"`
}

type PasswordChange struct {
	Password        string `form:"Password" binding:"required"`
	NewPassword     string `form:"NewPassword" binding:"required"`
	ConfirmPassword string `form:"ConfirmPassword" binding:"required"`
}
