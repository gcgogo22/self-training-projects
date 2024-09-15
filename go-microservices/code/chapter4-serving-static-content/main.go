package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	
	// http://localhost:3000/files/customer.csv 

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/servecontent", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./customers.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer customerFile.Close()

		http.ServeContent(w, r, "customerdata.csv", time.Now(), customerFile)
	})

	http.HandleFunc("/servefile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./customer.csv")
	})

}
