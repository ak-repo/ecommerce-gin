package authinterface

import (
	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	RegisterationHandler(ctx *gin.Context, role string)
	CustomerRegister(ctx *gin.Context)
	StoreRegister(ctx *gin.Context)
	DeliveryPartnerRegister(ctx *gin.Context)
	CustomerLogin(ctx *gin.Context)
	StoreLogin(ctx *gin.Context)
	DeliveryPartnerLogin(ctx *gin.Context)
	AdminLogin(ctx *gin.Context)
	AdminLoginForm(ctx *gin.Context)
	LoginHandler(ctx *gin.Context, role string)

	//password change
	CustomerPasswordChange(ctx *gin.Context)
	AdminPasswordChange(ctx *gin.Context)
	PasswordChangeHandler(ctx *gin.Context, role string)

	// otp
	SendOTPHandler(ctx *gin.Context)
	VerifyOTPHandler(ctx *gin.Context)
}

type AuthServiceInterface interface {
	RegisterService(input *auth.RegisterRequest, role string) error
	LoginService(input *auth.LoginRequest, role string) (*auth.LoginResponse, error)
	PasswordChangeService(userID uint, req *auth.PasswordChange) error

	//otp
	SentOTPService(req *auth.SendOTPRequest) error
	VerifyOTPService(req *auth.VerifyOTPRequest) error
}

type AuthRepoInterface interface {
	CreateUser(username, email, password, role string) error
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
