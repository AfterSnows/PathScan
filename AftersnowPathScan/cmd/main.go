package main

import (
	"AftersnowPathScan/core"
	"AftersnowPathScan/utils"
	"flag"
	"fmt"
)

func main() {

	var UrlFilename string
	var DictFilename string
	var ThreadNumber int
	var Url string
	fmt.Println("(-U=<targetUrl> or -UF=<target UrlFile> | -DF=<target UrlFile> [must have]|-T=<threads>) ")
	flag.StringVar(&UrlFilename, "UF", "", "filename默认为空")
	flag.StringVar(&DictFilename, "DF", "", "filename默认为空")
	flag.IntVar(&ThreadNumber, "T", 10, "线程默认为10")
	flag.StringVar(&Url, "U", "", "Url单独探测默认为空")
	flag.Parse()
	if Url != "" && UrlFilename == "" && DictFilename != "" {
		core.StartModel1(DictFilename, ThreadNumber, Url)
	} else if Url == "" && UrlFilename != "" && DictFilename != "" {
		URLS := utils.Open(UrlFilename)
		core.StartModel2(DictFilename, ThreadNumber, URLS)
	} else {
		fmt.Println("Please input according to the treaty of old King, may the sun guide your direction\n\n")

	}

}
