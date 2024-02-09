package router

import (
	//	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/book", BookRoutes{}.Routes())
	//r.Mount("/author", Author{}.Routes())
	r.Mount("/requests", RequestRoutes{}.Routes())

	return r
}
