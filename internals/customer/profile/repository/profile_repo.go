package profilerepo

import (
	"errors"

	profileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/profile/profile_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProfileRepo(db *gorm.DB) profileinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAddress(userID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("user_id=?", userID).First(&address).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &address, err
}

func (r *repository) UpdateAddress(address *models.Address) error {

	return r.DB.Save(address).Error
}

func (r *repository) AddAddress(address *models.Address) error {

	return r.DB.Create(address).Error
}
