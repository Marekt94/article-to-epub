package emailsender

import (
	"article-to-epub/pkg/modules"
	"fmt"
	"os"
	"sync"
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

	var emailSender modules.EmailSenderIntf = NewGmailEmailSender(user, pass)
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

	res := make(chan SendEmailResult, len(to))
	var wg sync.WaitGroup
	for i, receiver := range to {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			t.Logf("Worker number: %d", i)
			var emailSender modules.EmailSenderIntf = NewGmailEmailSender(user, pass)
			err := emailSender.SendEmail([]string{receiver}, "test topic", "test content", "test.html", []byte(file))
			res <- SendEmailResult{text: fmt.Sprintf("Worker %d: email to receiver %s sent with error %v", i, receiver, err), err: err}
		}(i)
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	for err := range res {
		t.Logf("%s", err.text)
		if err.err != nil {
			t.Fatalf("%v", err.err)
		}
	}
	t.Logf("Finished")
}
