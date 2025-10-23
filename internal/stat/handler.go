package stat

import (
	"net/http"
	"shorter-url/configs"
	"shorter-url/pkg/middleware"
	"shorter-url/pkg/response"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
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
		fromStr := r.URL.Query().Get("from")
		toStr := r.URL.Query().Get("to")

		if fromStr == "" {
			http.Error(w, "Missing 'from' parameter", http.StatusBadRequest)
			return
		}
		if toStr == "" {
			http.Error(w, "Missing 'to' parameter", http.StatusBadRequest)
			return
		}

		from, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			http.Error(w, "Invalid 'from' date: "+err.Error(), http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			http.Error(w, "Invalid 'to' date: "+err.Error(), http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		stats := handler.StatRepository.GetStats(by, from, to)
		response.JsonResponse(w, 200, stats)
	}
}
