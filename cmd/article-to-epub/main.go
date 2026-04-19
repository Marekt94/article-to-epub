package main

import (
	"article-to-epub/pkg/modules/articlesimplifier"
	"log"

	_ "github.com/Marekt94/go-kernel-mt"
)

func main() {
	simplifier := &articlesimplifier.ArticleSimplifier{}
	out, err := simplifier.SimplifyArticle([]byte("https://fs.blog/mental-models/?utm_source=unknownews"), 30)
	if err != nil {
		log.Fatalf("Error simplifying article: %v", err)
	}
	htmlToEpub := articlesimplifier.HtmlToEpubConverter{}
	_, err = htmlToEpub.ConvertHtmlToEpub(out)
	if err != nil {
		log.Fatalf("Error converting HTML to EPUB: %v", err)
	}
}
