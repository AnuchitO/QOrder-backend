package orders

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"gopkg.in/mgo.v2/bson"
)

func Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	order := Orders{}
	err := db.Orders().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = r.DecodeJsonPayload(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = db.Orders().Update(bson.M{"_id": order.ID}, order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
