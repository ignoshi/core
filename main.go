package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ignoshi/core/bookmarks"
	"github.com/ignoshi/core/snippets"
	"github.com/ignoshi/core/tags"
)

func main() {
	r := mux.NewRouter()
	tags.InjectRoutes(r.PathPrefix("/api/tags").Subrouter())
	bookmarks.InjectRoutes(r.PathPrefix("/api/bookmarks").Subrouter())
	snippets.InjectRoutes(r.PathPrefix("/api/snippets").Subrouter())

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
