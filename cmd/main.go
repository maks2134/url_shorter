package main

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/internal/auth"
	"shorter-url/internal/link"
	"shorter-url/internal/stat"
	"shorter-url/internal/user"
	"shorter-url/pkg/db"
	"shorter-url/pkg/event"
	"shorter-url/pkg/middleware"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepo := link.NewLinkRepository(database)
	userRepo := user.NewUserRepository(database)
	statRepo := stat.NewStatRepository(database)

	authService := auth.NewAuthService(userRepo)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepo,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepo, Config: conf, EventBus: eventBus})
	stat.NewStatHandler(router, stat.StatHandlerDeps{StatRepository: statRepo, Config: conf})

	go statService.AddClick()

	middlewareStack := middleware.Chain(middleware.CORS, middleware.Logging)

	return middlewareStack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Serve is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
