package profileinterface

import (
	adminprofile "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AdminProfileHandler(ctx *gin.Context)
	ShowAddressFormHandler(ctx *gin.Context)
	UpdateAddressHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	AdminProfileService(adminID uint) (*adminprofile.ProfileDTO, error)
	AdminAddressUpdateService(adminID, addressID uint, address *adminprofile.AddressDTO) error
	GetAdminAddressService(addressID uint) (*adminprofile.AddressDTO, error)
}

type RepoInterface interface {
	GetAdminAddressByUserID(userID uint) (*models.Address, error)
	GetAdminAddressByAddressID(AddressID uint) (*models.Address, error)
	AddAdminAdress(userID uint, address *adminprofile.AddressDTO) error
	UpdateAdminAddress(addressID uint, address *adminprofile.AddressDTO) error
}
