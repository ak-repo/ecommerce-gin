package profileservice

import (
	"errors"

	profiledto "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_dto"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_interface"
	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type service struct {
	ProfileRepo profileinterface.Repository
	UserRepo    authinterface.Repository
}

func NewProfileServiceMG(profileRepo profileinterface.Repository, userRepo authinterface.Repository) profileinterface.Service {
	return &service{ProfileRepo: profileRepo, UserRepo: userRepo}

}

func (s *service) GetProfile(userID uint) (*profiledto.ProfileDTO, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	address, err := s.ProfileRepo.GetAddress(userID)
	if err != nil || address == nil {
		return nil, errors.New("no address found")
	}

	profile := profiledto.ProfileDTO{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
		Role:  user.Role,
	}
	if address != nil {
		profile.Address = profiledto.AddressDTO{
			Phone:       address.Phone,
			ID:          address.ID,
			AddressLine: address.AddressLine,
			City:        address.City,
			State:       address.State,
			PostalCode:  address.PostalCode,
			Country:     address.Country,
		}
	}
	return &profile, nil
}

// add or update
func (s *service) UpdateAddress(input *profiledto.AddressDTO, userID uint) error {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	data, err := s.GetAddress(userID)
	if err != nil {
		return err
	}

	address := models.Address{
		UserID:      userID,
		AddressLine: input.AddressLine,
		City:        input.City,
		State:       input.State,
		PostalCode:  input.PostalCode,
		Country:     input.Country,
		Phone:       input.Phone,
	}
	if data == nil || data.ID == 0 {
		return s.ProfileRepo.AddAddress(&address)
	} else {
		address.Model = gorm.Model{ID: data.ID}
		return s.ProfileRepo.UpdateAddress(&address)
	}

}

// get address
func (s *service) GetAddress(userID uint) (*profiledto.AddressDTO, error) {

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	data, err := s.ProfileRepo.GetAddress(userID)
	if data == nil || err != nil {
		return nil, errors.New("no address found")
	}

	address := profiledto.AddressDTO{
		Phone:       data.Phone,
		ID:          data.ID,
		AddressLine: data.AddressLine,
		City:        data.City,
		State:       data.State,
		PostalCode:  data.PostalCode,
		Country:     data.Country,
	}

	return &address, nil
}
