package store

import "gopkg.in/mgo.v2/bson"

// Tag represents a tag model, which acts as a label for
// all other Ignoshi features.
// All Features in Ignoshi are built around tags
type Tag struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"name" bson:"description"`
}

// Save saves tag into database, either update it or create new one
func (t *Tag) Save() error {
	var err error
	if t.ID == bson.ObjectId("") {
		t.ID = bson.NewObjectId()
		err = db.C("tags").Insert(t)
	} else {
		err = db.C("tags").Update(bson.M{"_id": t.ID}, t)
	}
	return err
}

// ListTags returns a list of all tags exists in the db
func ListTags() ([]Tag, error) {
	tags := []Tag{}
	err := db.C("tags").Find(nil).All(&tags)
	return tags, err
}
