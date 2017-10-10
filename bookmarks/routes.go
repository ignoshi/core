package bookmarks

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emostafa/jopher"
	"github.com/gorilla/mux"
)

// InjectRoutes adds Tag domain routes to the main router
func InjectRoutes(r *mux.Router) {
	r.HandleFunc("/", listBookmarks).Methods("GET")
	r.HandleFunc("/", createBookmark).Methods("POST")
	r.HandleFunc("/{id}", getBookmark).Methods("GET")
}

func listBookmarks(w http.ResponseWriter, r *http.Request) {
	m := BookmarksManager{}
	items, err := m.Find(nil)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, items)
}

func createBookmark(w http.ResponseWriter, r *http.Request) {
	item := Bookmark{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		jopher.Error(w, 400, err)
		return
	}
	if ok, _ := item.IsValid(); !ok {
		jopher.Error(w, 400, errors.New("Invalid Body"))
		return
	}
	item.Save()
	jopher.Success(w, item)
}

func getBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := bson.M{"_id": vars["id"]}
	m := BookmarksManager{}
	item, err := m.FindOne(q)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, item)
}
