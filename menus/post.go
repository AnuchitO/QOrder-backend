package menus

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
)

func Post(w rest.ResponseWriter, r *rest.Request) {
	menu := Menus{
		ID: bson.NewObjectId(),
	}
	err := r.DecodeJsonPayload(&menu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	err = db.Menus().Insert(menu)
	if err != nil {
		fmt.Printf("% #v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
