package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/customers/add", func(w http.ResponseWriter, r *http.Request) {
		var c Customer
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&c)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Print(c)
	})

	go func() {
		time.Sleep(2 * time.Second)
		_, err := http.Post("http://localhost:3000/customers/add", "application/json", bytes.NewBuffer([]byte(
			`
				{
					"id": 999, 
					"firstName": "Arthur",
					"lastName": "Dent",
					"address": "155 Country Lane, Cottington, England"
				}
			`)))
		if err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := readCustomers()
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(customers)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

// Sending JSON message
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}

func convertToJSON(c Customer) ([]byte, error) {
	data, err := json.Marshal(c)
	return data, err
}

func convertToJSONEncoder(c Customer) ([]byte, error) {
	var b *bytes.Buffer
	enc := json.NewEncoder(b)
	// Encoder is re-usable, which allows the Encode function to be called multiple times.
	err := enc.Encode(c)
	return b.Bytes(), err
}

func convertFromJSON(data []byte) (Customer, error) {
	var c Customer
	err := json.Unmarshal(data, &c)
	return c, err
}

func convertFromJSONDecoder(data []byte) (Customer, error) {
	b := bytes.NewBuffer(data) // must be an io.Reader
	dec := json.NewDecoder(b)

	var c Customer
	// We can use this decoder multiple times in case we want to populate multiple objects
	err := dec.Decode(&c)
	return c, err
}

func readCustomers() ([]Customer, error) {
	f, err := os.Open("customers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	customers := make([]Customer, 0)
	csvReader := csv.NewReader(f)
	csvReader.Read() // throw away header

	for {
		_, err := csvReader.Read()
		if err == io.EOF {
			return customers, nil
		}
		if err != nil {
			return nil, err
		}

	}
}
