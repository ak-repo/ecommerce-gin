package usersrepository

import (
	"errors"

	userdto "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/user_dto"
	usersinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/user_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewrUsersRpository(db *gorm.DB) usersinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllUsers(req *userdto.UsersPagination) ([]models.User, error) {

	db := r.DB.Model(&models.User{})
	if req.Query != "" {
		db.Where("username ILIKE ?", "%"+req.Query+"%")
	}
	if req.Role != "" {
		db.Where("role=?", req.Role)
	}
	if req.Status != "" {
		db.Where("status=?", req.Status)
	}

	db.Count(&req.Total)

	offset := (req.Page - 1) * req.Limit

	var users []models.User
	err := db.Limit(req.Limit).
		Offset(offset).
		Find(&users).Error
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

func (r *repository) DeleteUser(userID uint) error {
	return r.DB.Delete(&models.User{}, userID).Error
}
