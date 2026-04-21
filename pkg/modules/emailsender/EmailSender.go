package emailsender

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	smtpHost string
	smtpPort int

	smtpUser string
	smtpPass string
}

func NewEmailSender(user string, password string) *EmailSender {
	return &EmailSender{
		smtpHost: "smtp.gmail.com",
		smtpPort: 587,
		smtpUser: user,
		smtpPass: password,
	}
}

func (es EmailSender) SendEmail(to string, topic string, content string, attachment []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", es.smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", topic)
	m.SetBody("text/plain", content)

	d := gomail.NewDialer(es.smtpHost, es.smtpPort, es.smtpUser, es.smtpPass)
	// Go requires either ServerName or InsecureSkipVerify when a custom tls.Config is provided.
	// ServerName should match the SMTP host for proper certificate verification.
	d.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12, ServerName: es.smtpHost}

	return d.DialAndSend(m)
}
