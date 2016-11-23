package orders

import (
	"github.com/ant0ine/go-json-rest/rest"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Get(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(findAllOrders())
}

func findAllOrders() (orders []Orders) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB("qorders").C("orders")
	c.Find(bson.M{}).All(&orders)
	return
}
