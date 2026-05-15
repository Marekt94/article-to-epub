package articlesimplifier

import (
	"bytes"
	"net/url"
	"os"

	log "github.com/Marekt94/go-kernel-mt/logging"
	"github.com/go-shiori/go-readability"
)

type ArticleSimplifierFormHTML struct {
}

func (a *ArticleSimplifierFormHTML) SimplifyArticle(article []byte) (simpArticle []byte, title string, author string, err error) {
	path := string(article)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", "", err
	}

	strm := bytes.NewReader(b)
	articleInt, err := readability.FromReader(strm, &url.URL{})
	if err != nil {
		return nil, "", "", err
	}
	log.Global.Infof("Article title: %s", articleInt.Title)
	log.Global.Tracef("Article content: %s", articleInt.Content)

	return []byte(articleInt.Content), articleInt.Title, articleInt.Byline, nil
}
