package service

import (
	"medical-app/src/model"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func EmailHandler(u model.User) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	mail := os.Getenv("EMAIL")
	pass := os.Getenv("EMAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", "Med-App@no-reply.com")
	m.SetHeader("To", u.Email)
	m.SetHeader("Subject", "Verify your account")
	m.SetBody("text/html", "Hello <b>"+u.Email+"</b>, <a href=http://localhost:3000/auth/verify?u="+u.Username+">click here</a> to verify your account!"+
		"\n"+
		"Thank you!")

	d := gomail.NewDialer("smtp.gmail.com", 587, mail, pass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
