package articlesimplifier

import (
	"os"
	"strconv"
	"time"

	log "github.com/Marekt94/go-kernel-mt/logging"
	"github.com/go-shiori/go-readability"
)

const defaultTimeout = 30

type ArticleSimplifierFromURL struct {
}

func (a *ArticleSimplifierFromURL) SimplifyArticle(article []byte) (simpArticle []byte, title string, authors string, err error) {
	var timeout int

	timeoutStr := os.Getenv("ARTICLE_SIMPLIFIER_TIMEOUT")

	if timeoutStr == "" {
		timeout = defaultTimeout
	} else if timeout, err = strconv.Atoi(timeoutStr); err != nil {
		log.Global.Warnf("Failed to parse timeout string: %v", err)
		timeout = defaultTimeout
	}

	out, title, authors, err := a.SimplifyArticleInt(article, timeout)
	return out, title, authors, err
}

func (a *ArticleSimplifierFromURL) SimplifyArticleInt(input []byte, timeout int) (simpArticle []byte, title string, authors string, err error) {
	url := string(input)
	article, err := readability.FromURL(url, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, "", "", err
	}
	log.Global.Infof("Article title: %s", article.Title)
	log.Global.Tracef("Article content: %s", article.Content)

	return []byte(article.Content), article.Title, article.Byline, nil
}
