package usersinterface

import (
	userdto "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/user_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	ChangeUserRole(ctx *gin.Context)
	BlockUser(ctx *gin.Context)
	ShowUserAddForm(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context) 
}

type Service interface {
	GetAllUsers(req *userdto.UsersPagination) ([]userdto.AdminUserListDTO, error)
	GetUserByID(userID uint) (*userdto.AdminUserDTO, error)
	ChangeUserRole(req *userdto.AdminUserRoleChange) error
	BlockUser(userID uint) error
	CreateUser(req *userdto.CreateUserRequest) (uint, error)
	DeleteUser(userID uint) error 
}

type Repository interface {
	GetAllUsers(req *userdto.UsersPagination) ([]models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	UpdateUser(user *models.User) error // generic
	CreateUser(user *models.User) error
	DeleteUser(userID uint) error 
}
