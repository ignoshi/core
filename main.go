package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ignoshi/core/tags"
)

func main() {
	r := mux.NewRouter()
	tags.InjectRoutes(r.PathPrefix("/api/tags").Subrouter())
	//r.HandleFunc("/api/tags", handlers.ListTags).Methods("GET")
	//r.HandleFunc("/api/tags", handlers.CreateTag).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
