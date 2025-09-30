package userrepository

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(username, email, password string) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserAddress(userID uint) (*models.Address, error)
	AddAddress(address *dto.AddressDTO, userID uint) error
	UpdateAddress(address *dto.AddressDTO) error
}

type userRepo struct {
	DB *gorm.DB
}

// Init auth repo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

// user registration
func (r *userRepo) CreateUser(username, email, password string) error {
	user := models.User{
		Email:        email,
		Username:     username,
		PasswordHash: password,
		Role:         "customer",
		IsActive:     true,
	}
	return r.DB.Create(&user).Error
}

// return user details
func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil 
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// return user address detials
func (r *userRepo) GetUserAddress(userID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("user_id=?", userID).First(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // return no error, just no address
	}
	if err != nil {
		return nil, err
	}

	return &address, nil
}

// update user address
func (r *userRepo) UpdateAddress(address *dto.AddressDTO) error {

	return r.DB.Model(&models.Address{}).Where("id=?", address.ID).Updates(map[string]interface{}{
		"address_line": address.AddressLine,
		"city":         address.City,
		"state":        address.State,
		"postal_code":  address.PostalCode,
		"country":      address.Country,
		"phone":        address.Phone,
	}).Error
}

func (r *userRepo) AddAddress(address *dto.AddressDTO, userID uint) error {

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
