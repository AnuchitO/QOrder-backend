package orders

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
)

func Post(w rest.ResponseWriter, r *rest.Request) {
	order := Orders{
		ID: bson.NewObjectId(),
	}
	err := r.DecodeJsonPayload(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = db.Orders().Insert(order)
	if err != nil {
		fmt.Printf("% #v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
