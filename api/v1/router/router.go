package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/user", User{}.Routes())
	r.Mount("/book", Book{}, Routes())
	r.Mount("/author", Author{}.Routes())
	r.Mount("/requests", Requests{}.Routes())

	return r
}
