package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// gorilla/mux, go-chi/chi

func main() {
	r := mux.NewRouter()

	// We can now encode our parameters directly within the handler.
	r.HandleFunc("/products/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Get parameters
		idRaw := vars["id"]
		if len(idRaw) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(idRaw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	http.Handle("/", r)

	/*
	go-chi // Have the http methods
	r := chi.NewRouter()

	r.Get("/products", func(w, r) {})
	idRaw := chi.URLParam(r, "id") // Use the key

	http.Handle("/", r)

	*/
}
