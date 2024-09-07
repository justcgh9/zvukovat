package services

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendActivationMail(to, link string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "zvukovat@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Confirm your email address")
	htmlBody := fmt.Sprintf(`
        <div>
            <h1>Good Afternoon!</h1>
            <p>To activate your account on zvukovat follow the link below.</p>
            <a href="%s">Zvukovat'!</a>
            <p>Have a good time using our service!</p>
        </div>
    `, link)
	m.SetBody("text/html", htmlBody)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
