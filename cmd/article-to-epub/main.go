package main

import (
	kernel "github.com/Marekt94/go-kernel-mt"
	l "github.com/Marekt94/go-kernel-mt/logging"
)

type ArticleToEpubKernel struct {
	kernel.Kernel
}

func (k *ArticleToEpubKernel) Init() {

}

func (k *ArticleToEpubKernel) RegisterModule(m kernel.ModuleIntf) {

}

func (k *ArticleToEpubKernel) Run() {

}

func (k *ArticleToEpubKernel) NewAtEKernel() kernel.KernelIntf {
	l.SetGlobalLogger(l.NewZerologLogger())
	l.Global.Infof("Logger initialized")
	l.Global.Infof("Creating new ArticleToEpubKernel instance")
	ke := ArticleToEpubKernel{Kernel: kernel.NewKernel()}
	l.Global.Infof("ArticleToEpubKernel instance created")
	return &ke
}

func main() {
	k := ArticleToEpubKernel{}
	k.Init()
	k.Run()
}
