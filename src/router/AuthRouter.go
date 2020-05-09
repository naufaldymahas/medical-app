package router

import (
	"medical-app/src/controller"

	"github.com/go-chi/chi"
)

func AuthRouter(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", controller.Login)
		r.Post("/register", controller.Register)
	})
}
