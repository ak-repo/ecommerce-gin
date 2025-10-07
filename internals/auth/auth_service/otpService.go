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

func (s *AuthService) SentOTPService(req *auth.SendOTPRequest) error {

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

func (s *AuthService) VerifyOTPService(req *auth.VerifyOTPRequest) error {
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

	return nil
}
