package userrepository

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type AdminUserRepo interface {
	AdminFindUserByID(userID uint) (*models.User, error)
	AdminGetAllUsers() ([]models.User, error)
}

type adminUserRepo struct {
	DB *gorm.DB
}

func NewAdminUserRepo(db *gorm.DB) AdminUserRepo {
	return &adminUserRepo{DB: db}
}

// user by id
func (r *adminUserRepo) AdminFindUserByID(userID uint) (*models.User, error) {
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
func (r *adminUserRepo) AdminGetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}
