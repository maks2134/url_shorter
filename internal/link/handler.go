package link

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
)

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{}

	router.HandleFunc("GET /{alias}", linkHandler.CreateLink)
	router.HandleFunc("POST /link", linkHandler.GetLink)
	router.HandleFunc("PATCH /link/{id}", linkHandler.UpdateLink)
	router.HandleFunc("DELETE /link/{id}", linkHandler.DeleteLink)
}

func (link *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Link")
}

func (link *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Link")
}

func (link *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Link")
}

func (link *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Link")
}
