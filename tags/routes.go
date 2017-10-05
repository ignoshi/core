package tags

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emostafa/jopher"
	"github.com/gorilla/mux"
)

// InjectRoutes adds Tag domain routes to the main router
func InjectRoutes(r *mux.Router) {
	r.HandleFunc("/", listTags).Methods("GET")
	r.HandleFunc("/", createTag).Methods("POST")
}

// listTags fetches list of tags from the database and returns
// a JSON response
func listTags(w http.ResponseWriter, r *http.Request) {
	tags, err := Find()
	if err != nil {
		jopher.Error(w, 500, err)
		return
	}
	jopher.Success(w, tags)
}

// CreateTag consumes the body request to create a new Tag
func createTag(w http.ResponseWriter, r *http.Request) {
	tag := New()
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	defer r.Body.Close()
	if ok, errs := tag.IsValid(); !ok {
		for _, e := range errs {
			log.Println(e)
		}
		jopher.Write(w, 400, errs)
		return
	}
	err = tag.Save()
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	jopher.Success(w, tag)
}
