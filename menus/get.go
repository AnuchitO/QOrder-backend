package menus

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"gopkg.in/mgo.v2/bson"
)

func Get(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(findAllMenus())
}

func findAllMenus() (menus []Menus) {
	db.Menus().Find(bson.M{}).All(&menus)
	return
}
