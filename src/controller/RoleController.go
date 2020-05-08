package controller

import (
	"medical-app/src/config"
	"medical-app/src/service"
	"net/http"
)

type RoleController struct {
	RoleService service.RoleService
}

var roleService = service.ProvideRoleService(config.InitDb())

func FindAllRole(w http.ResponseWriter, r *http.Request) {
	result := roleService.FindAll()
	response := map[string]interface{}{
		"statusCode": 200,
		"result":     result,
		"length":     len(result),
	}
	service.RespondWithJSON(w, 200, response)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	err := roleService.Create(r)

	if err != nil {
		service.ErrorHandler(w, err)
		return
	}
	response := map[string]interface{}{
		"statusCode": 200,
		"message":    "Role has been created!",
	}

	service.RespondWithJSON(w, 200, response)
}
