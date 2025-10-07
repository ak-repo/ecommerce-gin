package authservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
)

func (s *AuthService) PasswordChangeService(userID uint, req *auth.PasswordChange) error {

	if req.ConfirmPassword != req.NewPassword {
		return errors.New("confirm and new password not matching")
	}

	if req.Password == req.NewPassword {
		return errors.New("old password same as new password")
	}

	user, err := s.authRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if ok := utils.CompareHashAndPassword(req.Password, user.PasswordHash); !ok {
		return errors.New("incorrect password")
	}

	hashpassword, err := utils.HashPassword(req.ConfirmPassword)
	if err != nil {
		return err
	}

	return s.authRepo.PasswordChange(userID, hashpassword)

}
