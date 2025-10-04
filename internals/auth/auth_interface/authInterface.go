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
}

type AuthServiceInterface interface {
	RegisterService(input *auth.RegisterRequest, role string) error
	LoginService(input *auth.LoginRequest, role string) (*auth.LoginResponse, error)
}

type AuthRepoInterface interface {
	CreateUser(username, email, password, role string) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	UpdatePassword(userID uint, hashPassword string) error
}
