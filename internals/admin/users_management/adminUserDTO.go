package usersmanagement

import "time"

type AdminUserDTO struct {
	ID            uint       `json:"id"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	Role          string     `json:"role"`
	Status        string     `json:"status"`
	EmailVerified bool       `json:"email_verified"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type AdminUserListDTO struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Status        string `json:"status"`
	EmailVerified bool   `json:"email_verified"`
}

type AdminUserRoleChange struct {
	ID   uint   `json:"id" form:"-"`
	Role string `json:"role" form:"role" binding:"required"`
}

// Request DTO for creating a new admin user
type CreateUserRequest struct {
	Username        string `form:"username" binding:"required,min=3,max=255"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=8"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
	Role            string `form:"role" binding:"required,oneof=customer store admin"`
	Status          string `form:"status" binding:"required,oneof=active inactive"`
	EmailVerified   bool   `form:"-"`
}
