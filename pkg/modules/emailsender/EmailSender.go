package emailsender

import (
	"article-to-epub/pkg/modules"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"sync"

	"io"

	"github.com/Marekt94/go-kernel-mt/logging"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	smtpHost string
	smtpPort int

	smtpUser string
	smtpPass string
}

type EmailSendResult struct {
	receiver string
	err      error
}

func NewEmailSender(ident string) modules.EmailSenderIntf {
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	if user == "" {
		logging.Global.Panicf("SMTP_USER is empty")
		return nil
	}
	if pass == "" {
		logging.Global.Panicf("SMTP_PASS is empty")
		return nil
	}
	return NewGmailEmailSender(user, pass)
}

func NewGmailEmailSender(user string, password string) *EmailSender {
	return &EmailSender{
		smtpHost: "smtp.gmail.com",
		smtpPort: 587,
		smtpUser: user,
		smtpPass: password,
	}
}

func (es EmailSender) SendEmail(to []string, topic string, content string, attachmentName string, attachment []byte) error {
	return es.sendEmailsParallel(to, topic, content, attachmentName, attachment)
}

func (es EmailSender) sendEmail(to []string, topic string, content string, attachmentName string, attachment []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", es.smtpUser)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", topic)
	m.SetBody("text/plain", content)

	logging.Global.Infof("Sending email from: %s, to: %v, subject: %q, attachment: %q", es.smtpUser, to, topic, attachmentName)

	m.Attach(attachmentName, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachment)
		return err
	}))

	d := gomail.NewDialer(es.smtpHost, es.smtpPort, es.smtpUser, es.smtpPass)
	d.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12, ServerName: es.smtpHost}

	err := d.DialAndSend(m)
	if err == nil {
		logging.Global.Infof("Email sent successfully!")
	}
	return err
}

func (es EmailSender) sendEmailsParallel(to []string, topic string, content string, attachmentName string, attachment []byte) error {
	sendResult := make(chan EmailSendResult, len(to))
	var wg sync.WaitGroup
	for _, receiver := range to {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := es.sendEmail([]string{receiver}, topic, content, attachmentName, attachment)
			sendResult <- EmailSendResult{receiver: receiver, err: err}
		}()
	}
	go func() {
		wg.Wait()
		close(sendResult)
	}()
	var err error = nil
	for r := range sendResult {
		if r.err != nil {
			errText := fmt.Sprintf("Email to receiver %s sent with error %v", r.receiver, r.err)
			logging.Global.Errorf(errText)
			err = errors.Join(err, errors.New(errText))
		} else {
			logging.Global.Infof("Email to receiver %s sent successfully!", r.receiver)
		}
	}
	return err
}
