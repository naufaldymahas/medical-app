package service

import (
	"encoding/json"
	"medical-app/src/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

type RoleService struct {
	Db *gorm.DB
}

func ProvideRoleService(DB *gorm.DB) RoleService {
	return RoleService{
		Db: DB,
	}
}

func (rs RoleService) Create(r *http.Request) error {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	rs.Db.Create(&role)
	return err
}

func (rs RoleService) FindAll() []model.Role {
	var roles []model.Role
	rs.Db.Find(&roles).Debug()
	return roles
}
