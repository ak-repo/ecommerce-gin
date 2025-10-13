package profileinterface

import (
	profiledto "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetProfile(ctx *gin.Context)
	GetAddress(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
}

type Service interface {
	GetProfile(userID uint) (*profiledto.ProfileDTO, error)
	GetAddress(userID uint) (*profiledto.AddressDTO, error)
	UpdateAddress(address *profiledto.AddressDTO, userID uint) error
}

type Repository interface {
	GetAddress(userID uint) (*models.Address, error)
	UpdateAddress(address *models.Address) error
	AddAddress(address *models.Address) error
}
