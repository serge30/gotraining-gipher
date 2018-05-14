package main

import (
	"log"
	"net/http"

	"github.com/serge30/gotraining-gipher/api/web"
	"github.com/serge30/gotraining-gipher/storage"
)

func main() {
	store, err := storage.NewSqliteStorage("gifs.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer store.Close()

	router := web.GetRouters(store)

	log.Println("Listening on the port 8081...")
	err = http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalln(err)
	}
}
