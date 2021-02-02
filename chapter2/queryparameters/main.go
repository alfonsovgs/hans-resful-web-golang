package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//QueryHandler handles the given query parameters
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s!\n", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()
	r.UseEncodedPath()
	r.HandleFunc("/articles", QueryHandler)

	//We can get a dunamically generated API route by using the url methodç
	url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	fmt.Println(url.Path)
	fmt.Println(err.Error())

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
