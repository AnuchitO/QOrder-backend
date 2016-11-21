package main

import (
	"net/http/httptest"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/matryer/silk/runner"
)

func TestOrdersEndpoint(t *testing.T) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("qorders").C("orders")
	c.RemoveAll(bson.M{})
	s := httptest.NewServer(newAPI().MakeHandler())
	defer s.Close()

	//runner.New(t, s.URL).RunGlob(filepath.Glob("*.silk.md"))
	runner.New(t, s.URL).RunFile("./orders.silk.md")
}
