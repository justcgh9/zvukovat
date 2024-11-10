package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type Email struct {
    user string
    host string
    password string
    port int
}

func New(
    usr, host, psswd string,
    port int,
) Email {
    return Email{
        user: usr,
        host: host,
        password: psswd,
        port: port,
    }
}

func (e Email) SendEmail(to, link string) error {
    message := gomail.NewMessage()

	message.SetHeader("From", e.user)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Confirm your email address")
	htmlBody := fmt.Sprintf(`
        <div>
            <h1>Good Afternoon!</h1>
            <p>To activate your account on zvukovat follow the link below.</p>
            <a href="%s">Zvukovat'!</a>
            <p>Have a good time using our service!</p>
        </div>
    `, link)
	message.SetBody("text/html", htmlBody)

	dialer := gomail.NewDialer(e.host, e.port, e.user, e.password)
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
