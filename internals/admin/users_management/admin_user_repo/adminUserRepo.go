package adminuserrepo

import (
	"errors"
	"fmt"

	usersmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/users_management"
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

// user role change
func (r *AdminUserRepo) AdminUserRoleChange(user *usersmanagement.AdminUserRoleChange) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&models.User{}).Where("id=?", user.ID).Update("role", user.Role)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("no user found with id %d", user.ID)
		}
		return nil
	})
}

// user block and unblock
func (r *AdminUserRepo) AdminUserBlock(user *models.User) error {
	return r.DB.Save(user).Error
}

// create user
func (r *AdminUserRepo) AdminCreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
