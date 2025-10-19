package main

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/internal/auth"
	"shorter-url/internal/link"
	"shorter-url/internal/user"
	"shorter-url/pkg/db"
	"shorter-url/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	linkRepo := link.NewLinkRepository(database)
	userRepo := user.NewUserRepository(database)

	authService := auth.NewAuthService(userRepo)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepo, Config: conf})

	middlewareStack := middleware.Chain(middleware.CORS, middleware.Logging)

	server := http.Server{
		Addr:    ":8081",
		Handler: middlewareStack(router),
	}

	fmt.Println("Serve is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
