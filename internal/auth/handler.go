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
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := AuthHandler{
		Config: deps.Config,
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

		fmt.Println(payload)

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

		fmt.Println(payload)

		data := RegisterResponse{
			Token: authHandler.Config.Auth.Secret,
		}

		response.JsonResponse(w, http.StatusOK, data)
	}
}
