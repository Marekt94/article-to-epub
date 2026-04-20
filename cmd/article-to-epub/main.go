package main

import (
	a "article-to-epub/pkg/modules/articlesimplifier"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

const ERROR = `ERROR: %v`

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(ERROR, "Error loading .env file")
	}

	var artSimp a.ArticleSimplifierIntf
	var htmlToEpubController a.HtmlToEpubConverterIntf

	artSimp = &a.ArticleSimplifier{}
	htmlToEpubController = &a.HtmlToEpubConverter{}

	re := regexp.MustCompile(`[^\pL]+`)
	for {
		fmt.Println("Enter article address:")
		var url string
		fmt.Scanln(&url)

		html, title, authors, err := artSimp.SimplifyArticle([]byte(url))
		if err != nil {
			log.Printf(ERROR, err)
			continue
		}

		out, err := htmlToEpubController.ConvertHtmlToEpub(html, title, authors)
		if err != nil {
			log.Printf(ERROR, err)
			continue
		}

		url = strings.ReplaceAll(title, `\`, "")
		url = re.ReplaceAllString(title, "_")
		filePath := url + `.epub`
		err = os.WriteFile(filePath, out, 0644)
		if err != nil {
			log.Printf(ERROR, err)
			continue
		}

		log.Printf("File saved to: %s", filePath)
	}
}
