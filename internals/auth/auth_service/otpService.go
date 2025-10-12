package authservice

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
)

func (s *authService) SendOTP(req *auth.SendOTPRequest) error {

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}

	hash := sha256.Sum256([]byte(otp))
	hashHex := hex.EncodeToString(hash[:])

	expire := time.Now().Add(1 * time.Minute)

	otpRecord := models.EmailOTP{
		Email:     req.Email,
		CodeHash:  hashHex,
		ExpiresAt: expire,
		Used:      false,
	}

	if err := s.authRepo.CreateOTP(&otpRecord); err != nil {
		return err
	}

	// sent
	if err := utils.SendEmailWithSendGrid(req.Email, otp); err != nil {
		s.authRepo.DeleteOTP(&otpRecord)
		return err
	}
	return nil

}

func (s *authService) VerifyOTP(req *auth.VerifyOTPRequest) error {
	record, err := s.authRepo.VerifyOTP(req)
	if err != nil {
		return err
	}

	hash := sha256.Sum256([]byte(req.OTP))
	hashHex := hex.EncodeToString(hash[:])

	if record.CodeHash != hashHex {
		return errors.New("invalid otp")
	}
	record.Used = true
	s.authRepo.UpdateOTP(record)

	// set email verified true
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("verified email is not valid- no matching user is found")
	}
	if err := s.authRepo.UserEmailVerified(user.ID); err != nil {
		return err
	}

	return nil
}
