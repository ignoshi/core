package tags

import (
	"github.com/gorilla/mux"
)

// InjectRoutes adds Tag domain routes to the main router
func InjectRoutes(r *mux.Router) {
	r.HandleFunc("/", listTags).Methods("GET")
	r.HandleFunc("/", createTag).Methods("POST")
}
