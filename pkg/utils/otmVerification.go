package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Generate otp
func GenerateOTP(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invlaid otp length")
	}
	digits := "1234567890"
	otp := make([]byte, length)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		otp[i] = digits[int(otp[i])%len(digits)]
	}
	return string(otp), nil
}

// sendEmailWithSendGrid uses SENDGRID_API_KEY and FROM_EMAIL env vars
func SendEmailWithSendGrid(toEmail, otp string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	fromEmail := os.Getenv("FROM_EMAIL")

	println("api", apiKey)
	println("email", fromEmail)
	if apiKey == "" || fromEmail == "" {
		return errors.New("sendgrid configuration missing")
	}

	from := mail.NewEmail("FreshBox", fromEmail)
	subject := "Your OTP"
	to := mail.NewEmail("", toEmail)
	textContext := fmt.Sprintf("Your verification code: %s\nThis code will expire in 1 minutes.", otp)
	htmlContent := fmt.Sprintf("<p>Your verification code: <strong>%s</strong></p><p>This code will expire in 10 minutes.</p>", otp)

	message := mail.NewSingleEmail(from, subject, to, textContext, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	resp, err := client.Send(message)
	fmt.Println("SendGrid Status:", resp.StatusCode)
	fmt.Println("SendGrid Body:", resp.Body)
	fmt.Println("SendGrid Headers:", resp.Headers)
	return err

}
