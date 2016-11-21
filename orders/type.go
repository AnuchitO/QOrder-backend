package orders

import "gopkg.in/mgo.v2/bson"

type Orders struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	MenuID   bson.ObjectId `bson:"menu_id" json:"menu_id"`
	MenuName string        `bson:"menu_name" json:"menu_name"`
	Amount   int           `bson:"amount" json:"amount"`
	Price    float32       `bson:"price" json:"price"`
	Seat     string        `bson:"seat" json:"seat"`
	Status   string        `bson:"status" json:"status"`
}
