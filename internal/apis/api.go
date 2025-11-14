package apis

import (
	"encoding/json"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	"github.com/Antonious-Stewart/Aggregator/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks"`
}

func HealthHandler(conn db.Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		checks := map[string]string{}

		if err := conn.Ping(); err != nil {
			checks["database"] = "down"
		} else {
			checks["database"] = "up"
		}

		status := "ok"

		for _, v := range checks {
			if v == "down" {
				status = "degraded"
				break
			}
		}

		response := HealthResponse{
			Status: status,
			Checks: checks,
		}

		w.Header().Set("Content-Type", "application/json")
		if status == "degraded" {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(response)
	}
}

func Routes(pool *db.Database) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	version, err := config.GetVar("API_VERSION")

	if err != nil {
		log.Fatal(err)
	}

	basePath := "/api/" + version

	router.Get(basePath+"/health-check", HealthHandler(pool))

	return router
}
