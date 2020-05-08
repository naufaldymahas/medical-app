package router

import (
	"medical-app/src/controller"

	"github.com/go-chi/chi"
)

func RoleRouter(r *chi.Mux) {
	r.Route("/role", func(r chi.Router) {
		r.Get("/", controller.FindAllRole)
		r.Post("/", controller.CreateRole)
	})
}
