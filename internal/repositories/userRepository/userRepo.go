package userrepository

import (
	"strconv"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User, email string) error
	GetUserProfile(user *models.User) error
	UpdateUserProfile(id string, address *models.Address) error
	AddUserProfile(address *models.Address) error
}

type userRepo struct {
	DB *gorm.DB
}

// Init auth repo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

// user registration
func (r *userRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// // Check user email already register
// func (r *userRepo) GetUserByEmail(user *models.User) error {
// 	return r.DB.Where("email=?", user.Email).First(user).Error

// }

// // Get user profile
// func (r *userRepo) GetUserProfile(user *models.User) error {

// 	err := r.DB.Where("user_id = ?", user.ID).Find(&user.Addresses).Error
// 	if err != nil {
// 		return err
// 	}

// 	if len(user.Addresses) == 0 {
// 		// no addresses found
// 		return gorm.ErrRecordNotFound
// 	}

// 	return nil

// }

// func (r *userRepo) AddUserProfile(address *models.Address) error {

// 	return r.DB.Create(address).Error
// }

// func (r *userRepo) UpdateUserProfile(id string, address *models.Address) error {
// 	idUint, _ := strconv.ParseUint(id, 10, 32)
// 	return r.DB.Model(&models.Address{}).
// 		Where("id = ?", uint(idUint)).
// 		Select("*").
// 		Updates(address).Error
// }

func (r *userRepo) GetUserByEmail(user *models.User, email string) error {
	return r.DB.Where("email = ?", email).First(user).Error
}

func (r *userRepo) GetUserProfile(user *models.User) error {
	// Preload addresses
	return r.DB.Preload("Addresses").First(user, "id = ?", user.ID).Error
}

func (r *userRepo) AddUserProfile(address *models.Address) error {
	return r.DB.Create(address).Error
}

func (r *userRepo) UpdateUserProfile(id string, address *models.Address) error {
	idUint, _ := strconv.ParseUint(id, 10, 32)
	return r.DB.Model(&models.Address{}).
		Where("id = ?", uint(idUint)).
		Updates(map[string]interface{}{
			"address_line": address.AddressLine,
			"city":         address.City,
			"state":        address.State,
			"postal_code":  address.PostalCode,
			"country":      address.Country,
		}).Error
}
