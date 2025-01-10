package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello Request")

	// read the body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write response
	fmt.Fprintf(rw, "Hello %s", b)
}
