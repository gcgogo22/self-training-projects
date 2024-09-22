package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	logFile        = "./logo.jpg"
	maxContentType = 512
)

var allowedFileTypes = []string{
	"image/jpeg",
	"image/jgp",
	"image/gif",
	"image/png",
}

func main() {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buff := make([]byte, maxContentType)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileType := http.DetectContentType(buff)

	if isValidFileType(fileType) {
		fmt.Println(fileType)
	} else {
		fmt.Println("unknown file type uploaded")
	}
}

func isValidFileType(fileType string) bool {
	for _, allowedType := range allowedFileTypes {
		if fileType == allowedType {
			return true
		}
	}
	return false
}
