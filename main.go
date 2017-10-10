package main

import (
	"log"
	"net/http"
	"os"
	"time"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ignoshi/core/bookmarks"
	"github.com/ignoshi/core/tags"
)

func main() {
	r := mux.NewRouter()
	tags.InjectRoutes(r.PathPrefix("/api/tags").Subrouter())
	bookmarks.InjectRoutes(r.PathPrefix("/api/bookmarks").Subrouter())
	lr := gh.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Handler:      lr,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started, Listening on :8000")
	log.Fatal(srv.ListenAndServe())
}
