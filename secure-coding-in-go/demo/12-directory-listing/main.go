package main

import (
	"net/http"
)

func main() {
	// If you have a sensitive test directory, then using localhost:8080/test/ 
	// will enter into this directory. 

	// Go intend to find the index.html file, if fail to find it, it serves the entire directory
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
