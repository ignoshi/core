package core

import (
	"errors"

	"github.com/ignoshi/core/tags"

	"gopkg.in/mgo.v2/bson"
)

// BaseModel is a model that all other models (Except Tag) should inherit from
type BaseModel struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"title" bson:"body"`
	Tags  []tags.Tag    `json:"tags" body:"tags"`
}

// IsValid checks if model fields are valid before save
func (bm *BaseModel) IsValid() (ok bool, errs []error) {
	ok = true
	if bm.Title == "" {
		ok = false
		errs = append(errs, errors.New("Title is required"))
	}
	if bm.Body == "" {
		ok = false
		errs = append(errs, errors.New("Body is required"))
	}

	return
}
