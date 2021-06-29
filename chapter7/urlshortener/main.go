package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alfonsovgs/hands_web_service/chapter7/base62Example/base62"
	models "github.com/alfonsovgs/hands_web_service/chapter7/urlshortener/helper"
	"github.com/gorilla/mux"
)

type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}

	dbClient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create a new route
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbClient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbClient.GenerateShortURL).Methods("POST")

	port := ":8000"
	srv := &http.Server{
		Handler: r,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	// Get ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)

	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record

	postBody, _ := ioutil.ReadAll(r.Body)

	log.Fatalln(r)
	log.Fatalln(postBody)

	err := json.Unmarshal(postBody, &record)
	err = driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
	responseMap := map[string]string{"encoded_string": base62.ToBase62(id)}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}

}
