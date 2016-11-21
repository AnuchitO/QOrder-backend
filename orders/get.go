package orders

import (
	"github.com/ant0ine/go-json-rest/rest"
	"gopkg.in/mgo.v2/bson"
)

func Get(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(findOreders())
}

func findOreders() []Orders {
	return []Orders{
		Orders{
			ID:       bson.ObjectIdHex("57e3f8e85e812900039d4647"),
			MenuID:   bson.ObjectIdHex("58059ff853490a001277d862"),
			MenuName: "ต้มยำกุ้ง",
			Amount:   1,
			Price:    99,
			Seat:     "05",
			Status:   "Waiting",
		},
	}
}
