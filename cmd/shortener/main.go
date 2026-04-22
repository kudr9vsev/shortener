package main

import (
	"net/http"

	"github.com/kudr9vsev/shortener/internal/handler"
	"github.com/kudr9vsev/shortener/internal/repository"
	"github.com/kudr9vsev/shortener/internal/service"
)

func main() {

	linkRepo := repository.NewMemoryLinkRepository()
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", linkHandler.CreateLink)
	mux.HandleFunc("GET /{id}", linkHandler.GetLink)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
