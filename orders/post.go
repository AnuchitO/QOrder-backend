package orders

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ant0ine/go-json-rest/rest"
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

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB("qorders").C("orders")
	err = c.Insert(order)
	if err != nil {
		fmt.Printf("% #v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(map[string]string{"err_msg": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
