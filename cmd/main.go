package main

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/internal/auth"
	"shorter-url/internal/link"
	"shorter-url/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	linkRepo := link.NewLinkRepository(database)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepo})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}
