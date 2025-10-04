package customerprofilerepo

import (
	"errors"

	custprofile "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile"
	customerprofileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type CustomerProfileRepo struct {
	DB *gorm.DB
}

func NewCustomerProfileRepo(db *gorm.DB) customerprofileinterface.RepoInterface {
	return &CustomerProfileRepo{DB: db}
}

// return user address detials
func (r *CustomerProfileRepo) GetCustomerAddress(userID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("user_id=?", userID).First(&address).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // return no error, just no address
		}
		return nil, err
	}

	return &address, nil
}

// update user address
func (r *CustomerProfileRepo) UpdateAddress(address *custprofile.AddressDTO) error {

	return r.DB.Model(&models.Address{}).Where("id=?", address.ID).Updates(map[string]interface{}{
		"address_line": address.AddressLine,
		"city":         address.City,
		"state":        address.State,
		"postal_code":  address.PostalCode,
		"country":      address.Country,
		"phone":        address.Phone,
	}).Error
}

func (r *CustomerProfileRepo) AddAddress(address *custprofile.AddressDTO, userID uint) error {

	return r.DB.Create(&models.Address{
		UserID:      userID,
		AddressLine: address.AddressLine,
		City:        address.City,
		State:       address.State,
		PostalCode:  address.PostalCode,
		Country:     address.Country,
		Phone:       address.Phone,
	}).Error
}
