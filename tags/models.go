package tags

import (
	"errors"

	"github.com/ignoshi/core/db"
	"gopkg.in/mgo.v2/bson"
)

// Tag represents a tag model, which acts as a label for
// all other Ignoshi features.
// All Features in Ignoshi are built around tags
type Tag struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
}

// New creates a new Tag object
func New() Tag {
	return Tag{}
}

// Save saves tag into database, either update it or create new one
func (t *Tag) Save() error {
	db := db.GetDB()
	var err error
	if t.ID == bson.ObjectId("") {
		t.ID = bson.NewObjectId()
		err = db.C("tags").Insert(t)
	} else {
		err = db.C("tags").Update(bson.M{"_id": t.ID}, t)
	}
	return err
}

// IsValid checks if all tag fields are valid before save
func (t *Tag) IsValid() (ok bool, errs []error) {
	ok = true
	if t.Title == "" {
		ok = false
		errs = append(errs, errors.New("Title is required"))
	}
	if t.Body == "" {
		ok = false
		errs = append(errs, errors.New("Body is required"))
	}

	return
}

// Find returns a list of all tags exists in the db
func Find() ([]Tag, error) {
	db := db.GetDB()
	tags := []Tag{}
	err := db.C("tags").Find(nil).All(&tags)
	return tags, err
}
