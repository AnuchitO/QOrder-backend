package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/anuchitprasertsang/QOrder-backend/db"
	"github.com/anuchitprasertsang/QOrder-backend/orders"
)

var (
	allowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders = []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Query",
		"Language",
		"Origin",
	}
)

func main() {
	db.ConnectLocalHost()
	log.Fatal(http.ListenAndServe(":8888", newAPI().MakeHandler()))
}

func newAPI() (api *rest.Api) {
	router, err := rest.MakeRouter(
		rest.Get("/orders", orders.Get),
		rest.Post("/orders", orders.Post),
		rest.Put("/orders/:id", orders.Put),
	)

	if err != nil {
		log.Fatal(err)
	}

	api = rest.NewApi()
	api.Use(rest.DefaultCommonStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	api.SetApp(router)
	return
}
