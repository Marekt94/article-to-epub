package main

import "github.com/Marekt94/go-kernel-mt"

func main() {
	var k kernel.KernelIntf
	k = &ArticleToEpubKernel{Kernel: kernel.NewKernel()}
	k.Init()
	k.Run()
}
