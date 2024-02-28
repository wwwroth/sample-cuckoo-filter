package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	cuckoo "github.com/panmari/cuckoofilter"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Success bool `json:"success"`
	Data    struct {
		Key     string `json:"key"`
		Message string `json:"message"`
	} `json:"data"`
}

func startServer(cuckooFilter *cuckoo.Filter) {

	router := mux.NewRouter()

	router.HandleFunc("/key/{key}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key := vars["key"]

		msg := "Key not found in filter."
		if cuckooFilter.Lookup([]byte(key)) {
			msg = "Key found in filter."
		}

		response := Response{}
		response.Success = true
		response.Data.Key = key
		response.Data.Message = msg

		w.Header().Set("Content-Type", "application/json")

		// Marshal the struct into JSON and write it to response
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}).Methods("GET")

	router.HandleFunc("/key/{key}/remove", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key := vars["key"]

		cuckooFilter.Delete([]byte(key))

		response := Response{}
		response.Success = true
		response.Data.Key = key
		response.Data.Message = "Key deleted from filter."

		w.Header().Set("Content-Type", "application/json")

		// Marshal the struct into JSON and write it to response
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}).Methods("GET")

	router.HandleFunc("/key/{key}/insert", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key := vars["key"]

		cuckooFilter.Insert([]byte(key))

		response := Response{}
		response.Success = true
		response.Data.Key = key
		response.Data.Message = "Key inserted into filter."

		w.Header().Set("Content-Type", "application/json")

		// Marshal the struct into JSON and write it to response
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}).Methods("GET")

	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	msg := fmt.Sprintf("Starting server on port %s.", os.Getenv("HTTP_PORT"))
	log.Println(msg)

	err := http.ListenAndServe(port, router)
	if err != nil {
		panic("Unable to start server.")
	}

}
