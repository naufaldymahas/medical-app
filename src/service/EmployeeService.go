package service

import (
	"encoding/json"
	"fmt"
	"log"
	"medical-app/src/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

type EmployeeService struct {
	Db *gorm.DB
}

func ProvideEmployeeService(DB *gorm.DB) EmployeeService {
	return EmployeeService{
		Db: DB,
	}
}

func (es EmployeeService) Create(e model.Employee, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&e)
	hashPassword := hashAndSalt([]byte(e.Password))
	e.Password = hashPassword
	es.Db.Create(&e)
	return err
}

func (es EmployeeService) FindOne(r *http.Request) (model.Employee, bool) {
	var e model.Employee
	// var roles model.Role
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		log.Fatal(err)
	}
	plainPwd := []byte(e.Password)
	es.Db.Where("username = ?", e.Username).Preload("Roles").First(&e)
	isTrue := comparePassword(e.Password, plainPwd)
	fmt.Println("isTrue: ", isTrue)
	if isTrue == false {
		return e, false
	}
	return e, true
}

func (es EmployeeService) AddRole(r *http.Request) error {
	// var e model.Employee
	var er model.EmployeeRole
	err := json.NewDecoder(r.Body).Decode(&er)

	es.Db.Create(&er)

	return err
}
