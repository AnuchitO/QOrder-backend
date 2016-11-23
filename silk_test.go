package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/anuchitprasertsang/QOrder-backend/db"
	"github.com/anuchitprasertsang/QOrder-backend/menus"
	"github.com/anuchitprasertsang/QOrder-backend/orders"
	"github.com/matryer/silk/runner"
)

func TestOrdersEndpoint(t *testing.T) {
	db.ConnectLocalHost()
	c := db.Orders()
	c.RemoveAll(bson.M{})

	o := orders.Orders{
		ID:       bson.ObjectIdHex("57e3f8e85e812900039d4647"),
		MenuID:   bson.ObjectIdHex("58059ff853490a001277d862"),
		MenuName: "ต้มยำกุ้ง",
		Amount:   1,
		Price:    99,
		Seat:     "05",
		Status:   "Waiting",
	}

	err := c.Insert(o)
	if err != nil {
		fmt.Printf("% #v\n", err.Error())
		return
	}

	s := httptest.NewServer(newAPI().MakeHandler())
	defer s.Close()

	//runner.New(t, s.URL).RunGlob(filepath.Glob("*.silk.md"))
	runner.New(t, s.URL).RunFile("./orders.silk.md")
}

func TestMenusEndpoint(t *testing.T) {
	db.ConnectLocalHost()
	c := db.Menus()
	c.RemoveAll(bson.M{})

	m := menus.Menus{
		ID:       bson.ObjectIdHex("57e3f8e85e812900039b1234"),
		MenuName: "ต้มยำกุ้ง",
		Price:    99,
	}

	err := c.Insert(m)
	if err != nil {
		fmt.Printf("% #v\n", err.Error())
		return
	}

	s := httptest.NewServer(newAPI().MakeHandler())
	defer s.Close()

	//runner.New(t, s.URL).RunGlob(filepath.Glob("*.silk.md"))
	runner.New(t, s.URL).RunFile("./menus.silk.md")
}
