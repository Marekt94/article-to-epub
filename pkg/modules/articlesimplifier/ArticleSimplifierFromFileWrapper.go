package articlesimplifier

import (
	"article-to-epub/pkg/modules"
	"os"
)

type ArticleSimplifierFromFileWrapper struct {
	ArticleSimplifier modules.ArticleSimplifierIntf
}

func (w *ArticleSimplifierFromFileWrapper) SimplifyArticle(article []byte) (simpArticle []byte, title string, author string, err error) {
	path := string(article)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", "", err
	}

	return w.ArticleSimplifier.SimplifyArticle(b)
}
