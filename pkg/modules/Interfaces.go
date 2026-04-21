package modules

type ArticleSimplifierIntf interface {
	SimplifyArticle(article []byte) (simpArticle []byte, title string, author string, err error)
}

type HtmlToEpubConverterIntf interface {
	ConvertHtmlToEpub(htmlContent []byte, title string, authors string) ([]byte, error)
}

type EmailSenderIntf interface {
	SendEmail(to string, topic string, content string, attachment []byte) error
}
