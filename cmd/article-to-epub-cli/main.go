package main

import (
	"article-to-epub/pkg/modules"
	a "article-to-epub/pkg/modules/articlesimplifier"
	"article-to-epub/pkg/modules/emailsender"
	h "article-to-epub/pkg/modules/htmltoepubconverter"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	logging.SetGlobalLogger(logging.NewZerologLogger())
}

func main() {
	var artSimp modules.ArticleSimplifierIntf
	var htmlToEpubController modules.HtmlToEpubConverterIntf
	var emailSender modules.EmailSenderIntf

	Init()

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := os.Getenv("SMTP_TO")
	from := user

	if pass == "" {
		logging.Global.Warnf("WARNING: SMTP_PASS is empty")
	}

	if user == "" {
		logging.Global.Panicf("SMTP_USER is empty, skipping loop execution")
		return
	}

	if to == "" {
		logging.Global.Panicf("Receiver e-mail address is empty")
		return
	}

	artSimp = &a.ArticleSimplifier{}
	htmlToEpubController = &h.HtmlToEpubConverter{}
	emailSender = emailsender.NewGmailEmailSender(user, pass)

	re := regexp.MustCompile(`[^\pL]+`)
	for {
		fmt.Println("Enter article address:")
		var url string
		fmt.Scanln(&url)

		html, title, authors, err := artSimp.SimplifyArticle([]byte(url))
		if err != nil {
			logging.Global.Panicf("%v", err)
			continue
		}

		out, err := htmlToEpubController.ConvertHtmlToEpub(html, title, authors)
		if err != nil {
			logging.Global.Panicf("%v", err)
			continue
		}

		url = strings.ReplaceAll(title, `\`, "")
		url = re.ReplaceAllString(title, "_")
		filePath := url + `.epub`

		to := []string{to, from}
		content := "Article from article-to-epub converter"
		topic := "Article-to-epub converter: " + filePath
		attachmentName := filePath
		attachment := out
		err = emailSender.SendEmail(to, topic, content, attachmentName, attachment)
		if err != nil {
			logging.Global.Panicf("%v", err)
		} else {
			logging.Global.Infof("***Article sent successfully***")
		}
	}
}
