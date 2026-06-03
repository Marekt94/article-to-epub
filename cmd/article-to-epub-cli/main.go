package main

import (
	"article-to-epub/pkg/misc"
	"article-to-epub/pkg/modules"
	"article-to-epub/pkg/modules/articlesimplifier"
	"article-to-epub/pkg/modules/emailsender"
	"article-to-epub/pkg/modules/htmltoepubconverters/calibreconverter"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
)

const (
	InpUnknown int = iota
	InpURL
	InpHTML
)

var (
	cli Cli
)

type Cli struct {
	PathOrUrl  string   `arg:"" name:"pathOrUrl" help:"Enter URL or HTML file path for artricle to convert"`
	Email      []string `help:"To: email" env:"SMTP_TO" name:"email" short:"e" sep:","`
	SaveToFile string   `help:"Path to directory, where file will be saved" short:"f" type:"path"`
	NoSend     bool     `help:"Do not send article via email" env:"NO_SEND" name:"no-send" short:"n"`
}

func DetectInputType(pOu string) (int, error) {
	if f, err := os.Stat(pOu); (err == nil) && !f.IsDir() {
		ext := strings.ToLower(filepath.Ext(pOu))
		if ext == `.html` || ext == `.htm` {
			logging.Global.Infof("Input: HTML")
			return InpHTML, nil
		}
	}

	_, err := url.Parse(pOu)
	if err == nil {
		logging.Global.Infof("Input: URL")
		return InpURL, nil
	}

	return InpUnknown, err
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	logging.SetGlobalLogger(logging.NewZerologLogger())

	ctx := kong.Parse(&cli)
	_ = ctx
	logging.Global.Infof("url: %s, email: %s, path: %s, send: %v", cli.PathOrUrl, cli.Email, cli.SaveToFile, !cli.NoSend)
}

func main() {
	Init()

	articleName := misc.AdaptUrlToFileName(cli.PathOrUrl)
	controller := modules.ArticleToEpubController{}

	kind, err := DetectInputType(cli.PathOrUrl)
	if err != nil {
		logging.Global.Panicf("Invalid path or url: %v", cli.PathOrUrl)
		return
	}

	var articleSimplifier modules.ArticleSimplifierIntf
	var emailSender modules.EmailSenderIntf
	htmlConverter := &calibreconverter.HtmlToEpubConverter{}

	if len(cli.Email) == 0 || cli.NoSend {
		emailSender = nil
	} else {
		emailSender = emailsender.NewEmailSender("")
	}

	switch kind {
	case InpURL:
		articleSimplifier = &articlesimplifier.ArticleSimplifierFromURL{}
	case InpHTML:
		articleSimplifier = &articlesimplifier.ArticleSimplifierFromFileWrapper{
			ArticleSimplifier: &articlesimplifier.ArticleSimplifierFormHTML{},
		}
	}

	res, err := controller.ConvertArticle([]byte(cli.PathOrUrl), articleName, cli.Email,
		articleSimplifier, htmlConverter, emailSender)

	if err != nil {
		logging.Global.Panicf("%v", err.Error())
	}
	logging.Global.Infof("Attachment name: %s, sent via e-mail: %v, epub: %v", res.AttachmentName, res.SentByEmail, res.Epub != nil)

}
