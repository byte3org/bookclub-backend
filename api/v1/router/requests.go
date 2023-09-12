package router

import (
	requestshandler "github.com/byte3/bookclub/backend/api/v1/handlers"
	"github.com/byte3/bookclub/backend/api/v1/middlewares"
	"github.com/go-chi/chi"
)

type RequestRoutes struct{}

func (re RequestRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(resourceRequestRoutes chi.Router) {
        resourceRequestRoutes.Use(middlewares.RequestCtx)
        resourceRequestRoutes.Get("/", requestshandler.GetRequestDetails)
        resourceRequestRoutes.Delete("/", requestshandler.DeleteRequest)
	})

	r.Route("/", func(rootRequestRoutes chi.Router) {
		rootRequestRoutes.Get("/", requestshandler.GetAllRequests)
		rootRequestRoutes.Post("/", requestshandler.CreateRequest)
		rootRequestRoutes.Get("/count", requestshandler.GetAllRequestsCount)
		rootRequestRoutes.Get("/status/pending", requestshandler.GetAllPendingRequests)
		rootRequestRoutes.Get("/status/accepted", requestshandler.GetAllAcceptedRequests)
	})

	return r
}
