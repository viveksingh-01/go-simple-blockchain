package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishedOn string `json:"published_on"`
	ISBN string `json:"isbn"`
}

type BookCheckout struct {
	BookID string `json:"book_id"`
	User string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis bool `json:"is_genesis"`
}

func main() {
	const PORT = "5000"
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	fmt.Println("Server started at PORT:", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, r))
}