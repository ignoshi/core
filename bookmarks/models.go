package bookmarks

import (
	"github.com/ignoshi/core/core"
	"github.com/ignoshi/core/db"
	"gopkg.in/mgo.v2/bson"
)

// Bookmark represents a bookmark url
type Bookmark struct {
	core.BaseModel
	URL string `json:"url" bson:"url"`
}

type BookmarksManager struct{}

// Save saves a bookmark into database, either update it or create new one
func (b *Bookmark) Save() error {
	db := db.GetDB()
	var err error
	if b.ID == bson.ObjectId("") {
		b.ID = bson.NewObjectId()
		err = db.C("bookmarks").Insert(b)
	} else {
		err = db.C("bookmarks").Update(bson.M{"_id": b.ID}, b)
	}
	return err
}

// Find returns a list of all bookmarks
func (m *BookmarksManager) Find(queryParams bson.M) ([]Bookmark, error) {
	items := []Bookmark{}
	db := db.GetDB()
	err := db.C("bookmarks").Find(queryParams).All(&items)
	return items, err
}

func (m *BookmarksManager) FindOne(queryParams bson.M) (Bookmark, error) {
	item := Bookmark{}
	db := db.GetDB()
	err := db.C("bookmarks").Find(queryParams).One(&item)
	return item, err
}
