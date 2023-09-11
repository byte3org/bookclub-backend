package router

import (
	"github.com/byte3/bookclub/backend/api/v1/handlers/userhandler"
	"github.com/byte3/bookclub/backend/api/v1/middlewares"
	"github.com/byte3/bookclub/backend/helpers/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type User struct{}

func (u User) Routes() chi.Router {
	r := chi.NewRouter()

	// these endpoints are only available to admin service
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Use(middlewares.EnsureAdminAuth)
		// for now only admin authenticated users can get all the users
		r.Get("/", userhandler.GetAllUsers)
		r.Get("/count", userhandler.GetUserCount)
	})

	r.Group(func(r chi.Router) {
		r.Post("/register", userhandler.CreateUser)
		r.Post("/login", userhandler.AuthenticateUser)
	})

	// these endpoints are only available to authenticated users
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Use(middlewares.EnsureAuth)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.UserCtx)
			r.Get("/", userhandler.GetUserDetails)
			r.Delete("/", userhandler.DeleteUser)
		})
	})

	return r
}
