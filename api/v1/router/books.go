package router

import (
	bookshandler "github.com/byte3/bookclub/backend/api/v1/handlers"
	"github.com/byte3/bookclub/backend/api/v1/middlewares"
	"github.com/go-chi/chi"
)

type BookRoutes struct{}

func (b BookRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(bookRoutes chi.Router) {
		bookRoutes.Use(middlewares.RequestCtx)
		bookRoutes.Use(middlewares.EnsureAuth)
		bookRoutes.Get("/", bookshandler.GetBookDetails)
		// delete
		// get /status
		// patch /status
		// get /owner
	})

	r.Route("/", func(bookRoutes chi.Router) {
		bookRoutes.Use(middlewares.EnsureAuth)
		bookRoutes.Get("/", bookshandler.GetAllBooks)
		bookRoutes.Post("/", bookshandler.CreateBook)
		bookRoutes.Get("/available", bookshandler.GetAllAvailableBooks)
		bookRoutes.Route("/{name}", func(bookNameRoutes chi.Router) {
			bookNameRoutes.Use(middlewares.BookNameCtx)
			bookNameRoutes.Get("/", bookshandler.GetBooksByName)
			bookNameRoutes.Get("/available", bookshandler.GetAvailableBooksByName)
		})
	})

	return r
}
