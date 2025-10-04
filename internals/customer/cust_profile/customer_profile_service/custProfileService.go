package customerprofileservice

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	custprofile "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile"
	customerprofileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_interface"
)

type CustomerProfileService struct {
	CustomerProfileRepo customerprofileinterface.RepoInterface
	CustomerRepo        authinterface.AuthRepoInterface
}

func NewCustomerProfileService(customerProfileRepo customerprofileinterface.RepoInterface, customerRepo authinterface.AuthRepoInterface) customerprofileinterface.ServiceInterface {
	return &CustomerProfileService{CustomerProfileRepo: customerProfileRepo, CustomerRepo: customerRepo}

}

func (s *CustomerProfileService) CustomerProfileService(userID uint) (*custprofile.ProfileDTO, error) {
	user, err := s.CustomerRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	address, err := s.CustomerProfileRepo.GetCustomerAddress(user.ID)
	if err != nil {
		return nil, err
	}

	profile := custprofile.ProfileDTO{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
		Role:  user.Role,
	}
	if address != nil {
		profile.Address = custprofile.AddressDTO{
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
func (s *CustomerProfileService) CustomerAddressUpdateService(address *custprofile.AddressDTO, userID uint) error {
	user, err := s.CustomerRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	data, err := s.CustomerProfileRepo.GetCustomerAddress(userID)
	if err != nil {
		return err
	}
	if data == nil {
		return s.CustomerProfileRepo.AddAddress(address, user.ID)
	} else {
		address.ID = data.ID
		return s.CustomerProfileRepo.UpdateAddress(address)
	}

}

// get address
func (s *CustomerProfileService) CustomerAddressService(userID uint) (*custprofile.AddressDTO, error) {

	data, err := s.CustomerProfileRepo.GetCustomerAddress(userID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no address found")
	}

	address := custprofile.AddressDTO{
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
