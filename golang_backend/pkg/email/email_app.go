package email

import (
	"github.com/spf13/viper"
	"kron-x/internal/dto"
	"os"
)

func NewEmail() Email {
	return &email{
		MimeHeaders:       viper.GetString("smtp.mime_headers"),
		SenderEmail:       viper.GetString("smtp.sender_email"),
		SenderAppPassword: os.Getenv("APP_PASSWORD"),
		SmtpHost:          viper.GetString("smtp.host"),
		SmtpPort:          viper.GetString("smtp.port"),
		DirectorEmail:     viper.GetString("smtp.director_email"),
	}
}

type Email interface {
	SendEmailCode(email, code string) error
	SendEmailPassword(email, userPassword string) error
	SenEmailChangePassword(email, newPassword string) error
	SendOrderUr(emailData *dto.SendOrderEmailUr) error
	SendOrderPhyz(emailData *dto.SendOrderEmailPhyz) error
	SendFeedbackCall(emailData *dto.FeedbackCall) error
	SendFeedbackSelection(emailData *dto.FeedbackSelection) error
	SendSupport(emailData *dto.Support) error
}
