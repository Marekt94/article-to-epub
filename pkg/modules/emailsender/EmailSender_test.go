package emailsender

import (
	"article-to-epub/pkg/modules"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

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
	to := `marekt94@gmail.com`
	if user == "" || pass == "" || to == "" {
		t.Skip("set SMTP_USER / SMTP_PASS / SMTP_TO (or put them in .env.local) to run this test")
	}

	var emailSender modules.EmailSenderIntf = NewEmailSender(user, pass)
	err = emailSender.SendEmail(to, "test topic", "test content", nil)
	if err != nil {
		t.Fatalf("send failed: %v", err)
	}
}
