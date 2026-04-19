package articlesimplifier

import (
	"time"

	log "github.com/Marekt94/go-kernel-mt/logging"
	"github.com/go-shiori/go-readability"
)

type ArticleSimplifier struct {
}

func (a *ArticleSimplifier) SimplifyArticle(input []byte, timeout int) ([]byte, error) {
	url := string(input)
	article, err := readability.FromURL(url, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	log.Global.Infof("Article title: %s", article.Title)
	log.Global.Tracef("Article content: %s", article.Content)

	return []byte(article.Content), nil
}
