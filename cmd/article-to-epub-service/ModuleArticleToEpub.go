package main

import (
	"article-to-epub/pkg/misc"
	"article-to-epub/pkg/modules"
	"article-to-epub/pkg/modules/articlesimplifier"
	"article-to-epub/pkg/modules/emailsender"
	"article-to-epub/pkg/modules/htmltoepubconverter"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	topic = "Article-to-epub converter: %s"
	body  = "Article from article-to-epub converter"
)

type Request struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type Response struct {
	Sent bool `json:"sent"`
}

type ModuleArticleToEpub struct {
	server *gin.Engine
}

func (m *ModuleArticleToEpub) convertHtml(c *gin.Context) {

}

func (m *ModuleArticleToEpub) fetchUrl(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var a modules.ArticleSimplifierIntf

	a = &articlesimplifier.ArticleSimplifier{}
	html, title, author, err := a.SimplifyArticle([]byte(req.Url))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var h modules.HtmlToEpubConverterIntf

	h = &htmltoepubconverter.HtmlToEpubConverter{}
	var epub []byte
	epub, err = h.ConvertHtmlToEpub(html, title, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	attachmentName := misc.AdaptUrlToFileName(req.Url) + `.epub`

	if req.Email == "" {
		c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, attachmentName))
		c.Data(http.StatusOK, "application/epub+zip", epub)
		return
	}

	var e modules.EmailSenderIntf
	e = emailsender.NewEmailSender("")
	err = e.SendEmail([]string{req.Email},
		fmt.Sprintf(topic, attachmentName),
		body,
		attachmentName,
		epub,
	)

	if err == nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (m *ModuleArticleToEpub) ExposeMethods() {
	m.server.POST("/convert-html", m.convertHtml)
	m.server.POST("/fetch-url", m.fetchUrl)
}

func (m *ModuleArticleToEpub) RegisterPermissions() {

}

func (m *ModuleArticleToEpub) GetName() string {
	return "Article to EPUB"
}
