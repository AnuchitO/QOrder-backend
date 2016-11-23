package menus

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"gopkg.in/mgo.v2/bson"
)

func Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	menu := Menus{}
	err := db.Menus().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&menu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = r.DecodeJsonPayload(&menu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = db.Menus().Update(bson.M{"_id": menu.ID}, menu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
