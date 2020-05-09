package main

import (
	"fmt"
	"medical-app/src/config"
	"medical-app/src/model"
	"medical-app/src/router"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db := config.InitDb()
	db.Debug().AutoMigrate(&model.Employee{}, &model.User{}, &model.Role{})

	router.AuthRouter(r)
	router.EmployeeRouter(r)
	router.RoleRouter(r)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
