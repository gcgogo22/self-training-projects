package main

import (
	"log"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

// Modifies the Readdir behavior and returns the empty list. 
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func main() {
	/*
	When a request is made, http.FileServer looks for a file in the specified directory based on the URL path of the incoming request. 

	For example, if a request is made to /images/picture.jpg, the server looks for the file ./tmp/static/images/picture.jpg

	http.StripPrefix modifies the URL path before it is passed to the next handler (http.FileServer(fs) in this case). It removes the specified prefix from the URL path. In this case, the prefix is "/", meaning that the first / will be removed from the path of any incoming request before it's passed to the file server. 

	If a request is made to /index.html, StripPrefix("/", ...) will strip the leading "/", resulting in index.html. Then it looks for the file ./tmp/static/index.html.

	StripPrefix is a middleware. 

	Request: http://localhost:8080/images/picture.jpg
	StripPrefix: /images/picture.jpg -> images/picture.jpg
	FileServer: .tmp/static/images/picture.jpg

	Response: If the file exists, it's served to the client, otherwise, a 404 error is returned. 

	So, if try to view the directory, it's restricted
	localhost:8080/test/ returns 404 page not found. 
	*/

	fs := justFilesFilesystem{http.Dir("./tmp/static")}
	err := http.ListenAndServe(":8080", http.StripPrefix("/", http.FileServer(fs)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
