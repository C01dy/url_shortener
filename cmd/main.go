package main

import (
	"fmt"
	"net/http"
	"urlshort/router"
	"urlshort/storage"
	"urlshort/api"
)

func main() {

	// linkStorage := storage.NewMemoryStorage()

	linkStorage, err := storage.NewSqliteStorage("links.db")
	if err != nil {
		panic(err)
	}
	redirectHandler := api.RedirectHandler(linkStorage)
    createLinkHandler := api.CreateLinkHandler(linkStorage)

    r := router.NewRouter()
    r.Handle("/", redirectHandler)
    r.Handle("/api/v1/links", createLinkHandler)

    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", r)

}
