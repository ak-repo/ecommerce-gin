package customerprofileinterface

import (
	custprofile "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	CustomerProfileHandler(ctx *gin.Context)
	GetCustomerAddress(ctx *gin.Context)
	CustomerAddressUpdateHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	CustomerProfileService(userID uint) (*custprofile.ProfileDTO, error)
	CustomerAddressUpdateService(address *custprofile.AddressDTO, userID uint) error
	CustomerAddressService(userID uint) (*custprofile.AddressDTO, error)
}

type RepoInterface interface {
	GetCustomerAddress(userID uint) (*models.Address, error)
	UpdateAddress(address *custprofile.AddressDTO) error
	AddAddress(address *custprofile.AddressDTO, userID uint) error
}
