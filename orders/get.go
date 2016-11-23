package orders

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"gopkg.in/mgo.v2/bson"
)

func Get(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(findAllOrders())
}

func findAllOrders() (orders []Orders) {
	db.Orders().Find(bson.M{}).All(&orders)
	return
}
