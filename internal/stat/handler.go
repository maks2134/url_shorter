package stat

import (
	"fmt"
	"net/http"
	"shorter-url/configs"
	"shorter-url/pkg/middleware"
	"time"
)

const (
	FiltereByDay   = "day"
	FiltereByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		by := r.URL.Query().Get("by")
		if by != FiltereByDay && by != FiltereByMonth {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		fmt.Println(from, to, by)
	}
}
