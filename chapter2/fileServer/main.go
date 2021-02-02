package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	path, _ := os.Getwd()

	//Mapping to methos is posible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir(path+"/chapter2/fileserver/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
