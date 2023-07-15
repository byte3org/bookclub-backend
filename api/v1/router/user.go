package router

import "github.com/go-chi/chi"

type User struct{}

func (u User) Routes() chi.Router {
	r := chi.NewRouter()

	// these endpoints are only available to admin service
	r.Group(func(r chi.Router) {
		r.Use(ensureAdminAuth)
		// for now only admin authenticated users can get all the users
		r.Get("/", userhandler.GetAllUsers)
	})

	r.Post("/register", userhandler.CreateUser)
	r.Post("/login", userhandler.AuthenticateUser)

	// these endpoints are only available to authenticated users
	r.Group(func(r chi.Router) {
		r.Use(ensureAuth)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Get("/", userhandler.GetUserDetails)
			r.Delete("/", userhandler.DeleteUser)
		})
	})
}
