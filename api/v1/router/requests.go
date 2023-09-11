package router

import (
	requestshandler "github.com/byte3/bookclub/backend/api/v1/handlers"
	"github.com/byte3/bookclub/backend/api/v1/middlewares"
	"github.com/byte3/bookclub/backend/helpers/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type Request struct{}

func (re Request) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Get("/", requestshandler.GetAllRequests)
		r.Post("/", requestshandler.CreateRequest)
		r.Get("/count", requestshandler.GetAllRequestsCount)
		r.Get("/status/pending", requestshandler.GetAllPendingRequests)
		r.Get("/status/accepted", requestshandler.GetAllAcceptedRequests)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.RequestCtx)
			r.Get("/", requestshandler.GetRequestDetails)
			r.Get("/", requestshandler.DeleteRequest)
		})
	})

	return r
}
