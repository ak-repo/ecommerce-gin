package adminuserinterface

import (
	usersmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/users_management"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ListAllUsersHandler(ctx *gin.Context)
	ListUserByIDHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	AdminAllUsersService() ([]usersmanagement.AdminUserListDTO, error)
	AdminGetUserByIDService(userID uint) (*usersmanagement.AdminUserDTO, error)
}

type RepoInterface interface {
	AdminFindUserByID(userID uint) (*models.User, error)
	AdminGetAllUsers() ([]models.User, error)
}
