package main

import (
	"article-to-epub/pkg/misc"
	"article-to-epub/pkg/modules"
	"article-to-epub/pkg/modules/articlesimplifier"
	"article-to-epub/pkg/modules/emailsender"
	"article-to-epub/pkg/modules/htmltoepubconverters/calibreconverter"
	"article-to-epub/pkg/modules/htmltoepubconverters/gonejackconverter"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/gin-gonic/gin"
)

type RequestUrl struct {
	Url   string   `json:"url"`
	Email []string `json:"email"`
}

type ModuleArticleToEpub struct {
	server *gin.Engine
	apiKey string
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
		&articlesimplifier.ArticleSimplifierFromURL{},
		&gonejackconverter.HtmlToEpubConverter{},
		emailsender.NewEmailSender(""))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if !res.SentByEmail {
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

	receiverEmail := c.PostFormArray(`email`)
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
		&articlesimplifier.ArticleSimplifierFromURL{},
		&calibreconverter.HtmlToEpubConverter{},
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

func (m *ModuleArticleToEpub) Authorize(c *gin.Context) {
	h := c.GetHeader("Authorization")

	if !strings.HasPrefix(h, "API-Key") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	apiKey := strings.TrimSpace(strings.TrimPrefix(h, "API-Key"))

	if apiKey == "" || apiKey != m.apiKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func (m *ModuleArticleToEpub) ExposeMethods() {
	api := m.server.Group("/api")
	api.Use(m.Authorize)
	api.POST("/convert-html", m.convertHtml)
	api.POST("/fetch-url", m.fetchUrl)

	m.server.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, nil) })
}

func (m *ModuleArticleToEpub) RegisterPermissions() {

}

func (m *ModuleArticleToEpub) GetName() string {
	return "Article to EPUB"
}
