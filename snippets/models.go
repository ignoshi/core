package snippets

import (
	"github.com/ignoshi/core/core"
	"github.com/ignoshi/core/db"
	"gopkg.in/mgo.v2/bson"
)

// Snippet represents a code snipper model
type Snippet struct {
	core.BaseModel
}

type SnippetsManager struct{}

// Save saves a snippet into database, either update it or create new one
func (s *Snippet) Save() error {
	db := db.GetDB()
	var err error
	if s.ID == bson.ObjectId("") {
		s.ID = bson.NewObjectId()
		err = db.C("snippets").Insert(s)
	} else {
		err = db.C("snippets").Update(bson.M{"_id": s.ID}, s)
	}
	return err
}

// ListSnippets returns a list of all snippets exists in the db
func (m *SnippetsManager) Find(queryParams bson.M) ([]Snippet, error) {
	items := []Snippet{}
	db := db.GetDB()
	err := db.C("snippets").Find(queryParams).All(&items)
	return items, err
}

func (m *SnippetsManager) FindOne(queryParams bson.M) (Snippet, error) {
	item := Snippet{}
	db := db.GetDB()
	err := db.C("snippets").Find(queryParams).One(&item)
	return item, err
}
