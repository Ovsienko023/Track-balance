package router

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"

	"api/internal/core"
	"api/internal/interfaces/web/handlers"
)

func RegisterHTTPEndpoints(router chi.Router, c core.Core, fs *embed.FS) http.Handler {
	h := handlers.New(c, fs)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/echo", h.Echo)
		r.Get("/docs", h.Docs)

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
