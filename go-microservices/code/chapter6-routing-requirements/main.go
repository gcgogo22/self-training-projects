package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Pay attention to the single and multiple resources request pattern. /products vs. /products/...

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		/*
			Demo for string splitting: /products/3
			parts := strings.Split(r.URL.Path, "/")

			if len(parts) != 3 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(parts[2])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		*/

		/*
			Demo for using regular expression to handle parametric routing
			var pattern = `^\/customers\/(\d+?)\/address\/(\S+)`
			var exp regexp.Regexp = regexp.MustCompile(pattern)

			func handleFunc(w http.ResponseWriter, r *http.Request) {
				path := r.URL.Path // "/customers/123/address/city"

				matches := exp.FindStringSubmatch(path)
				// First element is always the full match.
				// ["/customers/123/address/city", "123", "city"]

				if len(matches) == 0 {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
		*/

		for _, p := range products {
			if p.ID == id {
				data, err := json.Marshal(p)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		return
	})
}

// Parametric Routing
// String Splitting

func handleFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // "/customers/123/address/city"

	// With string splitting
	parts := strings.Split(path, "/")
	// ["", "cusotmers", "123", "address", "city"]

	// One of the downside of string splitting is that we need to write quite a lot of validation logics to validate the parameters and route the request.
}
