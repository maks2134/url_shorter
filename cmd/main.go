package main

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/internal/auth"
	"shorter-url/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}
