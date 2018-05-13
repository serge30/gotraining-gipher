package main

import (
	"log"
	"net/http"

	"github.com/serge30/gotraining-gipher/api/web"
	"github.com/serge30/gotraining-gipher/storage"
)

func main() {
	store, err := storage.NewFakeStorage()
	if err != nil {
		log.Fatalln(err)
	}

	router := web.GetRouters(store)

	log.Println("Listening on the port 8081")
	err = http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalln(err)
	}
}
