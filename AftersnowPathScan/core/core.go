package core

import (
	"AftersnowPathScan/model"
	"AftersnowPathScan/utils"
	"fmt"
	"github.com/panjf2000/ants"
	"log"
	"sync"
)

var Fetch *ants.PoolWithFunc
var WG sync.WaitGroup
var Url string
var Result int

func StartModel1(filename string, ThreadNumber int, url string) {
	Start(filename, ThreadNumber, url)
	Wait()
	End()
}
func StartModel2(filename string, ThreadNumber int, urls []string) {
	for _, setUrl := range urls {
		Start(filename, ThreadNumber, setUrl)
		Wait()
		End()
	}

}
func Task(targetsMap interface{}) {
	defer WG.Done()
	Result = 0
	targets, ok := targetsMap.([]string)
	if !ok {
		return
	}
	UsefulPaths := make([]string, 0)
	for _, target := range targets {
		isFound := model.StatusCode(Url + target)
		if isFound == 200 {
			Result = 1
			UsefulPaths = append(UsefulPaths, target)
			fmt.Println(Url + target + " ====> FIND!")
			utils.WriteFile(Url, target)
		} else {
			fmt.Println(Url + target + " ====> SHIT!")
			continue
		}
	}
	if Result == 0 {
		fmt.Printf("Change another dict try try!")
	}
}

func Start(filename string, ThreadNumber int, setUrl string) {
	var err error
	SetUrl(setUrl)
	Fetch, err = ants.NewPoolWithFunc(ThreadNumber, Task)
	if err != nil {
		log.Fatal(err.Error())
	}
	ThreadTarget := utils.Open(filename)
	WG.Add(1)
	if err := Fetch.Invoke(ThreadTarget); err != nil {
		log.Fatal(err.Error())
	}
}

func SetUrl(url string) {
	Url = url
}

func Wait() {
	WG.Wait()
}

func End() {
	Fetch.Release()
}
