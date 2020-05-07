package service

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	if err1 := godotenv.Load(); err1 != nil {
		log.Fatal(err1)
	}

	salt, err2 := strconv.Atoi(os.Getenv("SALT"))
	if err2 != nil {
		log.Fatal(err2)
	}

	hash, err3 := bcrypt.GenerateFromPassword(pwd, salt)
	if err3 != nil {
		log.Fatal(err3)
	}

	return string(hash)
}

func comparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
