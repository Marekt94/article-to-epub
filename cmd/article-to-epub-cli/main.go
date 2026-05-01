package main

import (
	"log"
	"os"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
)

type CLI struct {
	Url        string `arg:"" name:"url" help:"Enter URL for artricle to convert"`
	Email      string `help:"To: email" name:"email" short:"e"`
	SaveToFile string `help:"Path to directory, where file will be saved" short:"f" type:"path"`
	NoSend     bool   `help:"Do not send article via email" name:"no-send" default:"false" short:"n"`
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	logging.SetGlobalLogger(logging.NewZerologLogger())
}

func main() {
	// var artSimp modules.ArticleSimplifierIntf
	// var htmlToEpubController modules.HtmlToEpubConverterIntf
	// var emailSender modules.EmailSenderIntf

	Init()

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := os.Getenv("SMTP_TO")
	// from := user

	var cli CLI
	ctx := kong.Parse(&cli)
	_ = ctx
	logging.Global.Infof("url: %s, email: %s, path: %s, send: %v", cli.Url, cli.Email, cli.SaveToFile, !cli.NoSend)

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

	// artSimp = &a.ArticleSimplifierFromURL{}
	// htmlToEpubController = &h.HtmlToEpubConverter{}
	// emailSender = emailsender.NewGmailEmailSender(user, pass)

	// re := regexp.MustCompile(`[^\pL]+`)
	// for {
	// 	fmt.Println("Enter article address:")
	// 	var url string
	// 	fmt.Scanln(&url)

	// 	html, title, authors, err := artSimp.SimplifyArticle([]byte(url))
	// 	if err != nil {
	// 		logging.Global.Panicf("%v", err)
	// 		continue
	// 	}

	// 	out, err := htmlToEpubController.ConvertHtmlToEpub(html, title, authors)
	// 	if err != nil {
	// 		logging.Global.Panicf("%v", err)
	// 		continue
	// 	}

	// 	url = strings.ReplaceAll(title, `\`, "")
	// 	url = re.ReplaceAllString(title, "_")
	// 	filePath := url + `.epub`

	// 	to := []string{to, from}
	// 	content := "Article from article-to-epub converter"
	// 	topic := "Article-to-epub converter: " + filePath
	// 	attachmentName := filePath
	// 	attachment := out
	// 	err = emailSender.SendEmail(to, topic, content, attachmentName, attachment)
	// 	if err != nil {
	// 		logging.Global.Panicf("%v", err)
	// 	} else {
	// 		logging.Global.Infof("***Article sent successfully***")
	// 	}
	// }
}
