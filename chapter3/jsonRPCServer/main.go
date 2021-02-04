package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args holds arguments passed to JSON-RPC service
type Args struct {
	ID string
}

// Book struct holds Book JSON structure
type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

// JSONServer struct
type JSONServer struct{}

// GiveBookDetail return a book mathced by Args
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	//Read JSON file and load data
	absPath, _ := filepath.Abs("chapter3/books.json")
	raw, readererr := ioutil.ReadFile(absPath)
	if readererr != nil {
		log.Println("error:", readererr)
		os.Exit(1)
	}

	//Unmarsgal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		os.Exit(1)
	}

	// Iterate over each book to find the fiven book
	for _, book := range books {
		if book.ID == args.ID {
			// If book found, fill reply with it
			*reply = book
			break
		}
	}

	return nil
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()

	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
