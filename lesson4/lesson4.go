package main

import (
	"lesson4/document_store"

)

func main() {
	store := document_store.NewStore()
}
	cfg := &document_store.CollectionConfig{PrimaryKey: "key"}
		created, collection := store.CreateCollection("store", cfg)
		if created {
			fmt.Println("created")
		} else {
			fmt.Println("not created")}
}