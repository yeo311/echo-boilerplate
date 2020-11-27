package main

import (
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/yeo311/echo-boilerplate/handler"
	"github.com/yeo311/echo-boilerplate/mongodb"
	"github.com/yeo311/echo-boilerplate/route"
	"github.com/yeo311/echo-boilerplate/store"
)

func main() {
	log.SetFlags(log.Lshortfile)
	r := route.Route()

	v1 := r.Group("/api")

	mongodb.Init()
	us := store.NewUserStore(mgm.CollectionByName("users"))
	h := handler.NewHandler(us)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
