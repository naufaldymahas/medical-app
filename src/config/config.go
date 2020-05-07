package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func InitDb() *gorm.DB {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal(err1)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	DBUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err2 := gorm.Open("postgres", DBUrl)
	// defer db.Close()

	if err2 != nil {
		panic(err2)
	}

	return db
}
