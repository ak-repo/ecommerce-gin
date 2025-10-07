package profilerepo

import (
	"errors"

	adminprofile "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type AdminProfileRepo struct {
	DB *gorm.DB
}

func NewAdminProfileRepo(db *gorm.DB) profileinterface.RepoInterface {
	return &AdminProfileRepo{DB: db}
}

// Get admin address
func (r *AdminProfileRepo) GetAdminAddressByUserID(userID uint) (*models.Address, error) {

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

// Get admin address
func (r *AdminProfileRepo) GetAdminAddressByAddressID(AddressID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("id=?", AddressID).First(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &address, nil
}

// Add address
func (r *AdminProfileRepo) AddAdminAdress(userID uint, address *adminprofile.AddressDTO) error {
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

func (r *AdminProfileRepo) UpdateAdminAddress(addressID uint, address *adminprofile.AddressDTO) error {

	return r.DB.Model(&models.Address{}).Where("id=?", addressID).Updates(map[string]interface{}{
		"address_line": address.AddressLine,
		"city":         address.City,
		"state":        address.State,
		"postal_code":  address.PostalCode,
		"country":      address.Country,
		"phone":        address.Phone,
	}).Error
}
