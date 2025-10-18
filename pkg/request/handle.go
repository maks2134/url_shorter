package request

import (
	"net/http"
	"shorter-url/pkg/response"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decode[T](r.Body)
	if err != nil {
		response.JsonResponse(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	err = IsValid[T](body)
	if err != nil {
		response.JsonResponse(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	return &body, nil
}
