package userrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type UserAuthRepo interface {
	CreateUser(user *models.User) error
}

type userAuthRepo struct {
	DB *gorm.DB
}

// Init auth repo
func NewUserAuthRepo(db *gorm.DB) UserAuthRepo {
	return &userAuthRepo{DB: db}
}

// user registration
func (r *userAuthRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
