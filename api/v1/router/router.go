package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("v1 endpoints"))
    })
//	r.Mount("/book", Book{}, Routes())
//	r.Mount("/author", Author{}.Routes())
	r.Mount("/requests", Request{}.Routes())

	return r
}
