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
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
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
	if t.Name == "" {
		ok = false
		errs = append(errs, errors.New("Name is required"))
	}
	if t.Description == "" {
		ok = false
		errs = append(errs, errors.New("Description is required"))
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
