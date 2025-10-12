package usersrepository

import (
	"errors"

	usersinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/user_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewrUsersRpository(db *gorm.DB) usersinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *repository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err

}

func (r *repository) UpdateUser(user *models.User) error {
	return r.DB.Model(&models.User{}).Where("id=?", user.ID).Updates(user).Error
}

func (r *repository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
