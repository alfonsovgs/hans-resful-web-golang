package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

func handle(w http.ResponseWriter, r *http.Request) {
	//Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currentlyin the check content type middleware")

		// Filtering requests by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsuported Media Type. Please send JSON"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)

		// Setting cookie to every API response
		cookie := http.Cookie{Name: "Server-Time (UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(originalHandler)

	http.Handle("city/", chain)
	http.ListenAndServe(":8000", nil)
}
