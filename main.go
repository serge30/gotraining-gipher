package main

import (
	"fmt"

	"github.com/serge30/gotraining-gipher/storage"
)

func main() {
	store, _ := storage.NewFakeStorage()
	items, _ := store.GetItems()

	for _, gif := range items {
		fmt.Println(gif)
	}
}
