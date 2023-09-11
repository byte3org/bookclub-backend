package router

import (
    "net/http"
	requestshandler "github.com/byte3/bookclub/backend/api/v1/handlers"
	"github.com/byte3/bookclub/backend/api/v1/middlewares"
	"github.com/go-chi/chi/v5"
)

type Request struct{}

func (re Request) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(rootRequestRoutes chi.Router) {
		rootRequestRoutes.Get("/", requestshandler.GetAllRequests)
		rootRequestRoutes.Post("/", requestshandler.CreateRequest)
		rootRequestRoutes.Get("/count", requestshandler.GetAllRequestsCount)
		rootRequestRoutes.Get("/status/pending", requestshandler.GetAllPendingRequests)
		rootRequestRoutes.Get("/status/accepted", requestshandler.GetAllAcceptedRequests)
	})

	r.Group(func(resourceRequestRoutes chi.Router) {
		resourceRequestRoutes.Route("/{id}", func(req  chi.Router) {
			req.Use(middlewares.RequestCtx)
			req.Get("/", requestshandler.GetRequestDetails)
			req.Delete("/", requestshandler.DeleteRequest)
		})
	})

	return r
}
