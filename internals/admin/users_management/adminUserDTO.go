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

// func ToAdminUserListDTO(u *models.User) *dto.AdminUserListDTO {
// 	return &AdminUserListDTO{
// 		ID:       u.ID,
// 		Username: u.Username,
// 		Email:    u.Email,
// 		Role:     u.Role,
// 		Status:   u.Status,
// 	}
// }

// func ToAdminUserDTO(u *models.User) *dto.AdminUserDTO {
// 	return &dto.AdminUserDTO{
// 		ID:            u.ID,
// 		Username:      u.Username,
// 		Email:         u.Email,
// 		Role:          u.Role,
// 		Status:        u.Status,
// 		EmailVerified: u.EmailVerified,
// 		LastLoginIP:   u.LastLoginIP,
// 		LastLoginAt:   u.LastLoginAt,
// 		CreatedAt:     u.CreatedAt,
// 		UpdatedAt:     u.UpdatedAt,
// 	}
// }
