package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"api/infrastructure/config"
	"api/internal/core"
	"api/internal/interfaces/web/handlers"
)

func RegisterHTTPEndpoints(router chi.Router, c core.Core, apiConfig *config.Api) http.Handler {
	h := handlers.New(c)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/echo", h.Echo)

		// USERS
		r.Get("/profile", h.GetProfile)

		// CIRCLES
		r.Get("/circle/{circle_id}", h.GetCircle)
		r.Get("/circles", h.SearchCircles)
		r.Post("/circle", h.CreateCircle)
		r.Delete("/circle/{circle_id}", h.DeleteCircle)
	})

	return router
}
