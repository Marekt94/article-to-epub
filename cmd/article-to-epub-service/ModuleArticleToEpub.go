package main

import (
	"article-to-epub/pkg/misc"
	"article-to-epub/pkg/modules"
	"article-to-epub/pkg/modules/articlesimplifier"
	"article-to-epub/pkg/modules/emailsender"
	"article-to-epub/pkg/modules/htmltoepubconverter"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/gin-gonic/gin"
)

type RequestUrl struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type ModuleArticleToEpub struct {
	server *gin.Engine
}

func (m *ModuleArticleToEpub) fetchUrl(c *gin.Context) {
	var req RequestUrl
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	articleName := misc.AdaptUrlToFileName(req.Url)

	controller := modules.ArticleToEpubController{}

	res, err := controller.ConvertArticle([]byte(req.Url), articleName, req.Email,
		&articlesimplifier.ArticleSimplifierFromURL{}, &htmltoepubconverter.HtmlToEpubConverter{},
		emailsender.NewEmailSender(""))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if res.SentByEmail {
		if (res.Epub != nil) && (res.AttachmentName != "") {
			c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, res.AttachmentName))
			c.Data(http.StatusOK, "application/epub+zip", res.Epub)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No .epub file nor attachment name"})
		}
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func (m *ModuleArticleToEpub) convertHtml(c *gin.Context) {
	fh, err := c.FormFile(`html`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	receiverEmail := c.PostForm(`email`)
	url := c.PostForm(`url`)
	if url == "" {
		url = "defaul name for article from article-to-epub software"
	}

	var f multipart.File
	f, err = fh.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	logging.Global.Infof(`File size: %v`, fh.Size)
	logging.Global.Infof(`Receiver email: %v`, receiverEmail)
	logging.Global.Infof(`Article name: %v`, url)

	var html []byte
	html, err = io.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	articleName := misc.AdaptUrlToFileName(url)
	controller := modules.ArticleToEpubController{}

	res, err := controller.ConvertArticle(html, articleName, receiverEmail,
		&articlesimplifier.ArticleSimplifierFromURL{}, &htmltoepubconverter.HtmlToEpubConverter{},
		emailsender.NewEmailSender(""))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if res.SentByEmail {
		if (res.Epub != nil) && (res.AttachmentName != "") {
			c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, res.AttachmentName))
			c.Data(http.StatusOK, "application/epub+zip", res.Epub)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No .epub file nor attachment name"})
		}
	} else {
		c.JSON(http.StatusOK, nil)
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
