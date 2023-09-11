package main

import (
	"log"
	"net/http"
	"time"

	"github.com/byte3/bookclub/backend/api/v1/router"
	"github.com/byte3/bookclub/backend/config"
	"github.com/byte3/bookclub/backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initializeBookclubRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("bookclub server"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", router.SetupRoutes())
	})
}

func initializeBookclubServer(config *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)

	r.Use(middleware.Timeout(20 * time.Second))

	initializeBookclubRoutes(r)

	return r
}

func main() {
	// load config file
	config := config.GetConfig()

	// initialize database
	database.Initialize(config)

	// initialize server
	r := initializeBookclubServer(config)

    log.Printf("[!] Starting bookclub services on %s:%s", config.Host, config.Port)

	// start listener
    http.ListenAndServe(config.Host+ ":" +config.Port, r)
}
