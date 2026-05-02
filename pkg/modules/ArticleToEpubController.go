package modules

import (
	"fmt"
)

const (
	topic = "Article-to-epub converter: %s"
	body  = "Article from article-to-epub converter"
)

type ArticleToEpubController struct {
}

type ConvertRes struct {
	Epub           []byte
	AttachmentName string
	SentByEmail    bool
}

func (m *ArticleToEpubController) ConvertArticle(
	article []byte,
	articleName string,
	receiverEmail string,
	a ArticleSimplifierIntf,
	h HtmlToEpubConverterIntf,
	e EmailSenderIntf) (*ConvertRes, error) {

	html, title, author, err := a.SimplifyArticle(article)
	if err != nil {
		return nil, err
	}

	var epub []byte
	epub, err = h.ConvertHtmlToEpub(html, title, author)
	if err != nil {
		return nil, err
	}

	attachmentName := articleName + `.epub`

	if receiverEmail == "" {
		return &ConvertRes{epub, attachmentName, false}, nil
	}

	err = e.SendEmail([]string{receiverEmail},
		fmt.Sprintf(topic, attachmentName),
		body,
		attachmentName,
		epub,
	)

	if err == nil {
		return &ConvertRes{epub, attachmentName, true}, nil
	} else {
		return nil, err
	}
}
