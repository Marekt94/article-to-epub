package main

import (
	"os"

	k "github.com/Marekt94/go-kernel-mt"
	"github.com/Marekt94/go-kernel-mt/logging"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ArticleToEpubKernel struct {
	k.Kernel

	server *gin.Engine
}

func (a *ArticleToEpubKernel) Init() {
	err := godotenv.Load()
	if err != nil {
		logging.Global.Panicf("error loading .env")
	}
	logging.SetGlobalLogger(logging.NewZerologLogger())

	logging.Global.Infof("Article-to-epub kernel initialization...")

	gin.DefaultWriter = logging.Global.Writer()
	gin.DefaultErrorWriter = logging.Global.Writer()

	a.server = gin.New()
	a.server.Use(gin.Recovery())
	a.server.Use(gin.Logger())

	a.RegisterModule(&ModuleArticleToEpub{a.server, os.Getenv("API_KEY")})

	logging.Global.Infof("Article-to-epub kernel initialization finished")
}

func (a *ArticleToEpubKernel) Run() {
	a.Kernel.Run()

	a.server.Run()
}
