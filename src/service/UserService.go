package service

import (
	"encoding/json"
	"log"
	"medical-app/src/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func ProvideUserService(DB *gorm.DB) UserService {
	return UserService{
		Db: DB,
	}
}

func (us UserService) Create(r *http.Request) (model.User, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	user.Password = hashAndSalt([]byte(user.Password))
	us.Db.Create(&user)

	return user, err
}

func (us UserService) FindOne(r *http.Request) (model.User, bool) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	if err := us.Db.Where("username = ?", user.Username).First(&user).Error; err != nil {
		return user, false
	}

	plainPwd := []byte(user.Password)
	isTrue := comparePassword(user.Password, plainPwd)
	if isTrue == false {
		return user, false
	}

	return user, true
}
