package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type MyMiddleware struct {
	Next http.Handler
}

func (m MyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// do things before next handler
	m.Next.ServeHTTP(w, r) // pass request to next handler
	// do things after next handler
}

/*
Resiger the middleware

s := http.Server{
	Addr: ":5000",
	Handler: &loggingMiddleware{next: carMux},
}
*/

type loggingMiddleware struct {
	next http.Handler
}

func (lm loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if lm.next == nil {
		// lm.next = cartMux
	}

	slog.Info(fmt.Sprintf("Received %v request on route: %v", r.Method, r.URL.Path))
	now := time.Now()

	lm.next.ServeHTTP(w, r)

	slog.Info(fmt.Sprintf("Response generated for %v request on route %v. Duration: %", r.Method, r.URL.Path, time.Since(now)))
}

// Route-specific middleware

// Register route-specific middleware

// Note cartsHandler is not a method on a type, so it doesn't satisfy the http.Handler interface and it's not named ServeHTTP.

// This works because http.HandlerFunc type statisfies the http.Handler interface. 

/*
type HandlerFunc func(w http.ResponseWriter, r *http.Request) 

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(wr, r)
}

You can use http.HandlerFunc to directly type-convert cartsHandler function, because http.HandlerFunc is a type that acts as an adapter, allowing a function with the signature func(w, r) to satisfies the http.Handler interface. 

var handler http.Handler = http.HandlerFunc(cartsHanlder)
*/

/*
2. Can You Use ServeHTTP(w, r) function directly as an http.Handler?

You can't use it directly as an http.Handler. The reason is **http.Handler is an interface that requires a **TYPE** to implement the ServeHTTP method, not just a function. **

It doesn't work because ServeHTTP is a function but not a type that implements ServeHTTP method. 

- Use http.HandlerFunc(funcInstance) to change this function into a HandlerFunc type satisfies Handler interface. 

- Create a Struct to implement http.Handler interface. 
*/

/*
A GET reqeuest retrieves both the header and the full body of the resource. 
A HEAD request retrieves only the headers without the body. 

By default, if you define a GET handler in Go, the Go HTTP server will automatically handle the HEAD request for the same endpoint by sending the same headers but without body. Unless you want to write cusom logic for HEAD requests with specific behavior.
*/

/*
Note if the middleware uses ReadAll method of the r.Body, then you should be careful about invoking the middleware for validation. 

If it's a GET request, then you don't need to invoke the middleware. Because now, we have an empty JSON input. 

var data map[string]interface{}
err := json.Unmarshal([]byte(""), &data)
if err != nil {
	fmt.Println("Error:", err)
}

if r.Method != http.MethodPost {
	vm.next.ServeHTTP(w, r)
	return 
}
*/
func createShoppingCartService() *http.Server {
	cartMux.Handle("/carts", &validationMiddleware{
		next: http.HandlerFunc(cartsHandler), // Do a type conversion.
	})

	s := http.Server{
		Addr:    ":5000",
		Handler: &loggingMiddleware{next: cartMux},
	}

	return &s
}

type validationMiddleware struct {
	next http.Handler
}

func (vm validationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if vm.next == nil {
		log.Panic("No next handler defined for validationMidleware")
	}

	// Read the body out. After the request body has been read. It's consumed and can't be read again unless it's reset.

	// In Go, r.Body is an io.ReadCloser, which means it can only be read once because it's a stream. When the body is read, it becomes empty (the stream is exhausted), so if you need to access it again later in the request handling, you have to reset it.
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var c Cart
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the customer ID to the customer service and make sure it's valid
	res, err := http.Head(fmt.Sprintf("http://localhost:3000/customers/%v", c.CustomerID))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode == http.StatusNotFound {
		log.Print("Invalid customer ID")
		w.WriteHeader(http.StatusBadRequest)
		return // Validation failure, request shouldn't pass on to the next handler.
	}

	// If it's valid, we will reset the request body and send to the next handler to handle.

	// r.Body expects an io.ReadCloser, which has both Read and Close methods. Since bytes.NewBuffer only implements Read, you wrap it in io.NopCloser. This utility adds a no-op (no operation) Close method to the buffer, effectively making it an io.ReadCloser. This is important because r.Body requires an io.ReadCloser.

	// The reason for resetting the request body (r.Body) is that the body is often read more than once during request handling, such as when:
	// You need to inspect or log the request body.
	// The request body needs to be passed to other parts of the code (e.g., deserialized into a struct).

	// You need io.NopCloser if you're resetting the request body with something that doesn't have a Close() method, like bytes.Buffer. It's necessary because r.Body expects an io.ReadCloser but bytes.NewBuffer only provides an io.Reader.

	b := bytes.NewBuffer(data)
	r.Body = io.NopCloser(b) //Byte buffer doesn't have close on it. Use decorator NopCloser -> add a no operation close function to the reader.

	vm.next.ServeHTTP(w, r)
}
