package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alfonsovgs/hands_web_service/chapter7/jsonstore/helper"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type DBClient struct {
	db *gorm.DB
}

type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

func main() {
	db, err := helper.InitBD()
	if err != nil {
		panic(err)
	}

	dbclient := &DBClient{db: db}
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/package", dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package", dbclient.GetPackagesByWeight).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

// PostPackage saves the package information
func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)

	Package.Data = string(postBody)
	driver.db.Save(&Package)

	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

// GetPackage fethes the original URL for the given
// encoded(short) string
func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)

	driver.db.First(&Package, vars["id"])
	var PackageData interface{}

	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

// GetPackagesByWeight fetches all packages with given weight
func (driver *DBClient) GetPackagesByWeight(w http.ResponseWriter, r *http.Request) {
	var packages []helper.Package
	weight := r.FormValue("weight")

	// Handle response details
	var query = "SELECT * FROM \"Package\" where data::json->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	w.Write(respJSON)
}
