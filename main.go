package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT = "5000"
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	fmt.Println("Server started at PORT:", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, r))
}