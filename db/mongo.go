package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var database string

func GetDatabase(name string) *mgo.Database {
	if session == nil {
		log.Fatal("You have to connect to db first!")
	}
	return session.DB(name)
}

func Connect(uri, dbname string) (err error) {
	if session != nil {
		return nil
	}
	session, err = mgo.Dial(uri)

	if err != nil {
		return err
	}

	session.SetMode(mgo.Strong, true)
	database = dbname
	return nil
}

func Collection(collection string) *mgo.Collection {
	return GetDatabase(database).C(collection)
}

func ConnectLocalHost() {
	uri := map[string]string{
		"URI": "mongodb://localhost:27017",
	}
	err := Connect(uri["URI"], "qorders")
	if err != nil {
		log.Fatal(err)
	}
}
