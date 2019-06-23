package main

import (
	"log"
	"net/http"

	"github.com/tonespy/simple-api/apis"
	"github.com/tonespy/simple-api/router"
)

func main() {

	userRoutes := apis.GenerateUserRoutes()
	router := router.NewRouter(userRoutes)

	addr := ":8540"
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
