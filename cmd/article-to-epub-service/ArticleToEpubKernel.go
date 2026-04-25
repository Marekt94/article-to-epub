package main

import (
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
	logging.SetGlobalLogger(logging.NewZerologLogger())

	logging.Global.Infof("Article-to-epub kernel initialization...")

	a.server = gin.Default()
	err := godotenv.Load()
	if err != nil {
		logging.Global.Panicf("error loading .env")
	}
	a.RegisterModule(&ModuleArticleToEpub{a.server})

	logging.Global.Infof("Article-to-epub kernel initialization finished")
}

func (a *ArticleToEpubKernel) Run() {
	a.Kernel.Run()

	a.server.Run()
}
