package controller

import (
	"medical-app/src/config"
	"medical-app/src/service"
	"net/http"
)

var authService = service.ProvideUserService(config.InitDb())

func Login(w http.ResponseWriter, r *http.Request) {
	u, err := authService.FindOne(r)
	if err != true {
		service.RespondWithJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"statusCode": http.StatusUnauthorized,
			"message":    "Wrong username or password!",
		})
		return
	}

	response := map[string]interface{}{
		"statusCode": http.StatusOK,
		"result": map[string]interface{}{
			"id":       u.ID,
			"username": u.Username,
			"fullname": u.Firstname + " " + u.Lastname,
			"phone":    u.Phone,
			"email":    u.Email,
		},
	}

	service.RespondWithJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	u, err := authService.Create(r)
	if err != nil {
		service.ErrorHandler(w, err)
		return
	}

	if err := service.EmailHandler(u); err != nil {
		service.ErrorHandler(w, err)
	}

	service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"statusCode": http.StatusOK,
		"message":    "User has been created!",
	})
}
