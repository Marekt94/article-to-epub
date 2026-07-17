package emailsender

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type SendEmailResult struct {
	text string
	err  error
}

func TestSendEmail_Integration(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("Workspace dir %v", wd)
	err := godotenv.Load(".env.local")
	if err != nil {
		t.Fatalf("env file not found: %v", err)
	}

	// Safety switch so tests don't send emails by accident.
	if os.Getenv("RUN_EMAIL_TESTS") != "1" {
		t.Skip("set RUN_EMAIL_TESTS=1 to run email integration test")
	}

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := []string{`marekt94@gmail.com`, `sjndksnkjnakjns@sdkjsndkj.pl`}
	file := "<html><body><h1>Test</h1></body></html>"
	if user == "" || pass == "" {
		t.Skip("set SMTP_USER / SMTP_PASS (or put them in .env.local) to run this test")
	}

	var emailSender EmailSender = *NewGmailEmailSender(user, pass)
	err = emailSender.SendEmail(to, "test topic", "test content", "test.html", []byte(file))
	if err != nil {
		t.Fatalf("send failed: %v", err)
	}
}

func TestSendEmail_SendEmailsParallel(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("Workspace dir %v", wd)
	err := godotenv.Load(".env.local")
	if err != nil {
		t.Fatalf("env file not found: %v", err)
	}

	// Safety switch so tests don't send emails by accident.
	if os.Getenv("RUN_EMAIL_TESTS") != "1" {
		t.Skip("set RUN_EMAIL_TESTS=1 to run email integration test")
	}

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := []string{`marekt94@gmail.com`, `marekt94@gmail.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`}
	file := "<html><body><h1>Test</h1></body></html>"
	if user == "" || pass == "" {
		t.Skip("set SMTP_USER / SMTP_PASS (or put them in .env.local) to run this test")
	}

	var emailSender EmailSender = *NewGmailEmailSender(user, pass)
	err = emailSender.sendEmailsParallel(to, "test topic", "test content", "test.html", []byte(file))
	if err != nil {
		t.Fatalf("send failed: %v", err)
	}
}

func TestSendEmail_SendEmailsOneByOne(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("Workspace dir %v", wd)
	err := godotenv.Load(".env.local")
	if err != nil {
		t.Fatalf("env file not found: %v", err)
	}

	// Safety switch so tests don't send emails by accident.
	if os.Getenv("RUN_EMAIL_TESTS") != "1" {
		t.Skip("set RUN_EMAIL_TESTS=1 to run email integration test")
	}

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := []string{`marekt94@gmail.com`, `marekt94@gmail.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`}
	file := "<html><body><h1>Test</h1></body></html>"
	if user == "" || pass == "" {
		t.Skip("set SMTP_USER / SMTP_PASS (or put them in .env.local) to run this test")
	}

	var emailSender EmailSender = *NewGmailEmailSender(user, pass)
	err = emailSender.sendEmail(to, "test topic", "test content", "test.html", []byte(file))
	if err != nil {
		t.Fatalf("send failed: %v", err)
	}
}

func TestSendEmail_SendEmailsParallelFail(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("Workspace dir %v", wd)
	err := godotenv.Load(".env.local")
	if err != nil {
		t.Fatalf("env file not found: %v", err)
	}

	// Safety switch so tests don't send emails by accident.
	if os.Getenv("RUN_EMAIL_TESTS") != "1" {
		t.Skip("set RUN_EMAIL_TESTS=1 to run email integration test")
	}

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := []string{`marekt94@gmail.com`, `marekt94@gmail.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`, `test@test623.com`}
	file := "<html><body><h1>Test</h1></body></html>"
	if user == "" || pass == "" {
		t.Skip("set SMTP_USER / SMTP_PASS (or put them in .env.local) to run this test")
	}

	var emailSender EmailSender = EmailSender{
		smtpHost: "smtp.gmail.com",
		smtpPort: 587,
		smtpUser: user,
		smtpPass: "",
	}
	err = emailSender.sendEmailsParallel(to, "test topic", "test content", "test.html", []byte(file))
	if err == nil {
		t.Fatalf("Should have send exception!")
	} else {
		t.Logf("Send exception: %v", err)
	}
}
