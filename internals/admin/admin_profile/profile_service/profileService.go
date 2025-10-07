package profileservice

import (
	"errors"

	adminprofile "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_interface"
	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
)

type AdminProfileService struct {
	AdminInfoRepo authinterface.AuthRepoInterface
	ProfileRepo   profileinterface.RepoInterface
}

func NewAdminProfileService(adminInfoRepo authinterface.AuthRepoInterface,
	profileRepo profileinterface.RepoInterface) profileinterface.ServiceInterface {
	return &AdminProfileService{AdminInfoRepo: adminInfoRepo, ProfileRepo: profileRepo}

}

// admin profile service
func (s *AdminProfileService) AdminProfileService(adminID uint) (*adminprofile.ProfileDTO, error) {

	admin, err := s.AdminInfoRepo.GetUserByID(adminID)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("user not found")
	}

	address, err := s.ProfileRepo.GetAdminAddressByUserID(admin.ID)
	if err != nil {
		return nil, err
	}

	profile := adminprofile.ProfileDTO{
		ID:    admin.ID,
		Name:  admin.Username,
		Email: admin.Email,
		Role:  admin.Role,
	}
	if address != nil {
		profile.Address =
			adminprofile.AddressDTO{
				ID:          address.ID,
				AddressLine: address.AddressLine,
				City:        address.City,
				State:       address.State,
				PostalCode:  address.PostalCode,
				Country:     address.Country,
				Phone:       address.Phone,
			}
	}

	return &profile, nil
}

// admin profile update or add
func (s *AdminProfileService) AdminAddressUpdateService(adminID, addressID uint, address *adminprofile.AddressDTO) error {

	admin, err := s.AdminInfoRepo.GetUserByID(adminID)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("user not found")
	}

	if addressID == 0 {

		return s.ProfileRepo.AddAdminAdress(admin.ID, address)

	} else {
		return s.ProfileRepo.UpdateAdminAddress(addressID, address)
	}
}

// Get admin addrese
func (s *AdminProfileService) GetAdminAddressService(addressID uint) (*adminprofile.AddressDTO, error) {

	data, err := s.ProfileRepo.GetAdminAddressByAddressID(addressID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no address found")
	}

	address := adminprofile.AddressDTO{
		ID:          data.ID,
		AddressLine: data.AddressLine,
		City:        data.City,
		State:       data.State,
		PostalCode:  data.PostalCode,
		Country:     data.Country,
		Phone:       data.Phone,
	}
	return &address, nil
}
