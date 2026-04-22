package main

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/kudr9vsev/shortener/internal/config"
	"github.com/kudr9vsev/shortener/internal/handler"
	"github.com/kudr9vsev/shortener/internal/repository"
	"github.com/kudr9vsev/shortener/internal/service"
	"github.com/lib/pq"
)

func main() {
	config := db.LoadDBConfig()

	cfg, err := pq.NewConfig(config.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := sql.OpenDB(c)
	defer db.Close()

	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", linkHandler.CreateLink)
	mux.HandleFunc("GET /{id}", linkHandler.GetLink)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
