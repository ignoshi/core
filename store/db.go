package store

import mgo "gopkg.in/mgo.v2"

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("ignoshi")
}

// GetDB returns object of the database connection
func GetDB() *mgo.Database {
	return db
}
