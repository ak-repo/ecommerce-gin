package adminrepository

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type AdminRepo interface {
	GetAdminInfo(email string) (*models.User, error)
	GetAdminAddress(userID uint) (*models.Address, error)
	AddAdminAdress(userID uint, address *dto.AddressDTO) error
	UpdateAdminAddress(addressID uint, address *dto.AddressDTO) error
}

type adminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &adminRepo{DB: db}
}

// Get admin data from db
func (r *adminRepo) GetAdminInfo(email string) (*models.User, error) {
	var admin models.User
	err := r.DB.Where("email=?", email).First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, err
}

// Get admin address
func (r *adminRepo) GetAdminAddress(userID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("user_id=?", userID).First(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &address, nil
}

// Add address
func (r *adminRepo) AddAdminAdress(userID uint, address *dto.AddressDTO) error {
	return r.DB.Create(&models.Address{
		UserID:      userID,
		Phone:       address.Phone,
		AddressLine: address.AddressLine,
		State:       address.State,
		City:        address.City,
		PostalCode:  address.PostalCode,
		Country:     address.Country,
	}).Error
}

func (r *adminRepo) UpdateAdminAddress(addressID uint, address *dto.AddressDTO) error {

	return r.DB.Model(&models.Address{}).Where("id=?", addressID).Updates(map[string]interface{}{
		"address_line": address.AddressLine,
		"city":         address.City,
		"state":        address.State,
		"postal_code":  address.PostalCode,
		"country":      address.Country,
		"phone":        address.Phone,
	}).Error
}
