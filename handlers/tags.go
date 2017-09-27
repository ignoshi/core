package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emostafa/jopher"
	"github.com/ignoshi/core/store"
)

// ListTags fetches list of tags from the database and returns
// a JSON response
func ListTags(w http.ResponseWriter, r *http.Request) {
	tags, err := store.ListTags()
	if err != nil {
		jopher.Error(w, 500, err)
		return
	}
	jopher.Success(w, tags)
}

// CreateTag consumes the body request to create a new Tag
func CreateTag(w http.ResponseWriter, r *http.Request) {
	tag := store.Tag{}
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	defer r.Body.Close()

	err = tag.Save()
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	jopher.Success(w, tag)
}
