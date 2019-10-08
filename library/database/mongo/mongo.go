package mongo

import (
	"gopkg.in/mgo.v2"
)

type DB struct {
	*mgo.Database
}

func (d *DB) Copy(dbname ...string) *DB {
	if len(dbname) != 0 {
		return &DB{d.Session.Copy().DB(dbname[0])}
	}
	return &DB{d.Session.Copy().DB("")}
}

func (d *DB) Close() {
	d.Session.Close()
}

func NewDb(uri string) (db *DB) {
	if session, err := mgo.Dial(uri); err != nil {
		panic(err)
	} else {
		db = new(DB)
		//session.SetPoolLimit()
		db.Database = session.DB("")
	}
	return
}
