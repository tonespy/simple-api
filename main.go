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

	log.Fatal(http.ListenAndServe(":8540", router))
}
