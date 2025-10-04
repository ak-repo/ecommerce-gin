package adminuserrepo

import (
	"errors"

	adminuserinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type AdminUserRepo struct {
	DB *gorm.DB
}

func NewAdminUserRepo(db *gorm.DB) adminuserinterface.RepoInterface {
	return &AdminUserRepo{DB: db}
}

// user by id
func (r *AdminUserRepo) AdminFindUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

// all users
func (r *AdminUserRepo) AdminGetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}
