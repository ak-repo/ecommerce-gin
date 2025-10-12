package authinterface

import (
	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Registeration(ctx *gin.Context, role string)
	CustomerRegister(ctx *gin.Context)
	CustomerLogin(ctx *gin.Context)
	AdminLogin(ctx *gin.Context)
	AdminLoginForm(ctx *gin.Context)
	Login(ctx *gin.Context, role string)
	AdminLogout(ctx *gin.Context)
	Logout(ctx *gin.Context, role string)

	//password change
	CustomerPasswordChange(ctx *gin.Context)
	AdminPasswordChange(ctx *gin.Context)
	PasswordChange(ctx *gin.Context, role string)

	// otp
	SendOTP(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
}

type Service interface {
	 Registeration(input *auth.RegisterRequest, role string) error
	 Login(input *auth.LoginRequest, role string) (*auth.LoginResponse, error) 
	PasswordChange(userID uint, req *auth.PasswordChange) error

	//otp
	SendOTP(req *auth.SendOTPRequest) error
	VerifyOTP(req *auth.VerifyOTPRequest) error
}

type Repository interface {
	Registeration(user *models.User) error 
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	PasswordChange(userID uint, password string) error

	//otp
	CreateOTP(record *models.EmailOTP) error
	DeleteOTP(record *models.EmailOTP) error
	VerifyOTP(req *auth.VerifyOTPRequest) (*models.EmailOTP, error)
	UpdateOTP(record *models.EmailOTP) error
	UserEmailVerified(userID uint) error
}
