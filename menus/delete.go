package menus

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"gopkg.in/mgo.v2/bson"
)

func Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	err := db.Menus().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
