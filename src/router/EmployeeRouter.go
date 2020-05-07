package router

import (
	"medical-app/src/controller"

	"github.com/go-chi/chi"
)

func EmployeeRouter(r *chi.Mux) {
	r.Route("/employee", func(r chi.Router) {
		r.Post("/find", controller.FindOne)
		r.Post("/make", controller.Create)
	})
}
