package adminuserinterface

import (
	usersmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/users_management"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ListAllUsersHandler(ctx *gin.Context)
	ListUserByIDHandler(ctx *gin.Context)
	AdminUserRoleChangeHandler(ctx *gin.Context)
	AdminUserBlockHandler(ctx *gin.Context)
	AdminUserAddFormShowHandler(ctx *gin.Context)
	AdminUserCreationHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	AdminAllUsersService(query string) ([]usersmanagement.AdminUserListDTO, error)
	AdminGetUserByIDService(userID uint) (*usersmanagement.AdminUserDTO, error)
	AdminUserRoleChangeService(user *usersmanagement.AdminUserRoleChange) error
	AdminUserBlockService(userID uint) error
	AdminUserCreateService(req *usersmanagement.CreateUserRequest) (uint, error)
}

type RepoInterface interface {
	AdminFindUserByID(userID uint) (*models.User, error)
	AdminGetAllUsers() ([]models.User, error)
	AdminUserRoleChange(user *usersmanagement.AdminUserRoleChange) error
	AdminUserBlock(user *models.User) error
	AdminCreateUser(user *models.User) error
}
