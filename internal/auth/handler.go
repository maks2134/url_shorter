package auth

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/pkg/request"
	"shorter-url/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())

}

func (authHandler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		email, err := authHandler.AuthService.Login(payload.Email, payload.Password)
		fmt.Println(email, err)

		data := LoginResponse{
			Token: authHandler.Config.Auth.Secret,
		}

		response.JsonResponse(w, http.StatusOK, data)
	}
}

func (authHandler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		authHandler.AuthService.Register(payload.Email, payload.Password, payload.Name)
	}
}
