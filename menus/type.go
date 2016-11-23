package menus

import "gopkg.in/mgo.v2/bson"

type Menus struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	MenuName string        `bson:"menu_name" json:"menu_name"`
	Price    float32       `bson:"price" json:"price"`
}
