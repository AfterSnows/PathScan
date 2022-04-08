package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Open(filename string) (paths []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	return paths
}

func WriteFile(url string, value string) {
	var filepath string
	filepath = "UsefulPaths" + ".txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(url + value + "\n")
	write.Flush()
}
