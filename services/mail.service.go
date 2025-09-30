package services 
import (
	"log"
	"gopkg.in/gomail.v2"
)

type MailService struct{}

func NewMailService() *MailService{
	return &MailService{}
}
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "mikhael <mikhaeljhon22@gmail.com>"
const CONFIG_AUTH_EMAIL = "mikhaeljhon22@gmail.com"
const CONFIG_AUTH_PASSWORD = "hwzx uxlk drra ucsm"


func (s *MailService) SendEmail(to, subject, body string) error {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", CONFIG_SENDER_NAME)
    mailer.SetHeader("To", to)
    mailer.SetHeader("Subject", subject)
    mailer.SetBody("text/html", body)

    dialer := gomail.NewDialer(
        CONFIG_SMTP_HOST,
        CONFIG_SMTP_PORT,
        CONFIG_AUTH_EMAIL,
        CONFIG_AUTH_PASSWORD,
    )

    err := dialer.DialAndSend(mailer)
    if err != nil {
        log.Println("Failed to send email:", err)
        return err
    }

    log.Println("Mail sent!")
    return nil
}