package articlesimplifier

type ArticleSimplifierIntf interface {
	SimplifyArticle(article []byte) ([]byte, error)
}

type HtmlToEpubConverterIntf interface {
	ConvertHtmlToEpub(htmlContent []byte) ([]byte, error)
}
