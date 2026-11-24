package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		fmt.Println(r.Method)
		fmt.Println(r.URL.Query())
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Couldn't read request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to read request body"))
			return
		}
		fmt.Println(string(body))
		w.Write([]byte("Received request"))
	})
	fmt.Println("Server has started...")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
