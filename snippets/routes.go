package snippets

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
	r.HandleFunc("/", listSnippets).Methods("GET")
	r.HandleFunc("/", createSnippet).Methods("POST")
	r.HandleFunc("/{id}", getSnippet).Methods("GET")
}

func listSnippets(w http.ResponseWriter, r *http.Request) {
	m := SnippetsManager{}
	items, err := m.Find(nil)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, items)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	item := Snippet{}
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

func getSnippet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := bson.M{"_id": vars["id"]}
	m := SnippetsManager{}
	item, err := m.FindOne(q)
	if err != nil {
		jopher.NotFound(w, err)
		return
	}
	jopher.Success(w, item)
}
