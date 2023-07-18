package main

import (
	"log"
	"net/http"
	"time"

	"github.com/byte3/bookclub/backend/api/v1/router"
	"github.com/byte3/bookclub/backend/config"
	"github.com/byte3/bookclub/backend/helpers/jwt"
	"github.com/byte3/bookclub/backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initializeBookclubRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("destra command and control"))
	})
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/", router.SetupRoutes())
	})
}

func initializeBookclubServer(config *config.Config) *chi.Mux {
	r := chi.NewRouter()

	auth := jwt.JWT{}.New()
	r.Use(
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
		auth.Verifier(),
	)

	r.Use(middleware.Timeout(20 * time.Second))

	initializeBookclubRoutes(r)

	return r
}

func main() {
	log.Println("[!] Starting bookclub services...")

	// load config file
	config := config.GetConfig()

	// initialize database
	database.Initialize(config)

	r := initializeBookclubServer(config)

}
