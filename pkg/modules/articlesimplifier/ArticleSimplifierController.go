package articlesimplifier

import (
	"os"
	"strconv"
	"time"

	log "github.com/Marekt94/go-kernel-mt/logging"
	"github.com/go-shiori/go-readability"
)

const defaultTimeout = 30

type ArticleSimplifier struct {
}

func (a *ArticleSimplifier) SimplifyArticle(article []byte) ([]byte, error) {
	var timeout int
	var err error

	timeoutStr := os.Getenv("ARTICLE_SIMPLIFIER_TIMEOUT")

	if timeoutStr == "" {
		timeout = defaultTimeout
	} else if timeout, err = strconv.Atoi(timeoutStr); err != nil {
		log.Global.Warnf("Failed to parse timeout string: %v", err)
		timeout = defaultTimeout
	}

	out, err := a.SimplifyArticleInt(article, timeout)
	return out, err
}

func (a *ArticleSimplifier) SimplifyArticleInt(input []byte, timeout int) ([]byte, error) {
	url := string(input)
	article, err := readability.FromURL(url, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	log.Global.Infof("Article title: %s", article.Title)
	log.Global.Tracef("Article content: %s", article.Content)

	return []byte(article.Content), nil
}
