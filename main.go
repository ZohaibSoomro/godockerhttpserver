package main

import (
	"log"
	h "net/http"
)

func main() {
	h.HandleFunc("/hi", func(w h.ResponseWriter, r *h.Request) {
		w.Write([]byte("Hi boss!"))
	})

	h.HandleFunc("/", func(w h.ResponseWriter, r *h.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(h.ListenAndServe(":8080", nil))
}
