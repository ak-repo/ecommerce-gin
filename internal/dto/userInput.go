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

type LoginResponse struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *models.User
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
	PostalCode string `json:"zip_code"`
	Country    string `json:"country"`
}
