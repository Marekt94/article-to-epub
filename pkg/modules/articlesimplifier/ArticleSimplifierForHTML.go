package articlesimplifier

import (
	"bytes"
	"net/url"

	log "github.com/Marekt94/go-kernel-mt/logging"
	"github.com/go-shiori/go-readability"
)

type ArticleSimplifierFormHTML struct {
}

func (a *ArticleSimplifierFormHTML) SimplifyArticle(article []byte) (simpArticle []byte, title string, author string, err error) {
	//DONE: - to wywalić, bo to powinno być w main dla CLI (testy nie przechodza)
	strm := bytes.NewReader(article)
	articleInt, err := readability.FromReader(strm, &url.URL{})
	if err != nil {
		return nil, "", "", err
	}
	log.Global.Infof("Article title: %s", articleInt.Title)
	log.Global.Tracef("Article content: %s", articleInt.Content)

	return []byte(articleInt.Content), articleInt.Title, articleInt.Byline, nil
}
