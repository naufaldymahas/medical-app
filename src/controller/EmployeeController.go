package controller

import (
	"medical-app/src/config"
	"medical-app/src/model"
	"medical-app/src/service"
	"net/http"
)

var employeeService = service.ProvideEmployeeService(config.InitDb())

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	err := employeeService.Create(employee, r)
	if err != nil {
		service.ErrorHandler(w, err)
		return
	}

	response := map[string]interface{}{
		"statusCode": 200,
		"message":    "Employee has been created!",
	}

	service.RespondWithJSON(w, http.StatusOK, response)
}

func FindOneEmployee(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"statusCode": 200,
	}

	e, isTrue := employeeService.FindOne(r)

	if isTrue == false {
		response["statusCode"] = 400
		response["message"] = "Unauthenticated!"
		service.RespondWithJSON(w, 400, response)
		return
	}

	response["result"] = map[string]interface{}{
		"username":  e.Username,
		"firstname": e.Firstname,
		"lastname":  e.Lastname,
		"email":     e.Email,
		"roles":     e.Roles,
	}

	service.RespondWithJSON(w, 200, response)
}

func AddRoleToEmployee(w http.ResponseWriter, r *http.Request) {
	err := employeeService.AddRole(r)
	if err != nil {
		service.ErrorHandler(w, err)
		return
	}

	response := map[string]interface{}{
		"statusCode": 200,
		"message":    "Success adding role to employee",
	}

	service.RespondWithJSON(w, 200, response)
}
