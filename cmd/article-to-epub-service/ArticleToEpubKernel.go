package main

import (
	"article-to-epub/pkg/misc"
	"os"
	"path/filepath"
	"strings"

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
	// Load optional .env (local/dev convenience). Missing file shouldn't crash prod.
	if err := godotenv.Load(); err != nil {
		// Logger isn't set up yet, so keep it silent here.
	}
	logging.SetGlobalLogger(logging.NewZerologLoggerWithGinWritter())

	logging.Global.Infof("Article-to-epub kernel initialization...")

	gin.DefaultWriter = logging.Global.Writer()
	gin.DefaultErrorWriter = logging.Global.Writer()

	appMode := os.Getenv("MODE")
	if strings.ToUpper(appMode) != "RELEASE" {
		appMode = "DEBUG"
	}
	logging.Global.Infof("Application mode: %s", appMode)

	ginMode := os.Getenv("GIN_MODE")
	if strings.ToLower(ginMode) == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	a.server = gin.New()
	a.server.Use(gin.Recovery())
	a.server.Use(gin.Logger())

	appDir, err := misc.GetAppDir()
	if err != nil {
		logging.Global.Panicf("Application directory: %s", err.Error())
	} else {
		logging.Global.Infof("Application directory: %s", filepath.Dir(appDir))
	}

	a.RegisterModule(&ModuleArticleToEpub{a.server, os.Getenv("API_KEY")})

	logging.Global.Infof("Article-to-epub kernel initialization finished")
}

func (a *ArticleToEpubKernel) Run() {
	a.Kernel.Run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := a.server.Run(":" + port); err != nil {
		logging.Global.Panicf("gin server failed to start: %v", err)
	}
}
