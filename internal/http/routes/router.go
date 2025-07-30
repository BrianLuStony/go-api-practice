package routes

import (
	"example/api/internal/http/handlers"
	"example/api/internal/http/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(userHandler *handlers.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.Recover)
	r.Use(middlewares.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signin", userHandler.SignIn)
		r.Group(func(r chi.Router) {
			// r.Use(middlewares.JWTAuth)

			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.GetUser)
				r.Get("/{id}", userHandler.GetUserWithId)

			})
		})
	})

	return r

}
