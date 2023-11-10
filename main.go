package main

import (
	"event-management-service/config"
	"event-management-service/router"
	"log"

	"net/http"
)

func main() {

	eventDB, err := infra.DBCon()
	if err != nil {
		log.Fatal(err)
	}

	routes := router.Route(eventDB)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	errS := server.ListenAndServe()
	log.Fatal(errS)
}
