package Mail

import (
	"net/smtp"
	"strings"
)

type Mailer struct {
	Host     string
	User     string
	Password string
	To       string
	Title    string
	Content  string
}

func Send(host, user, password, to, title, content string) error {
	auth := smtp.PlainAuth("", user, password, host)
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" + content)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func (mail *Mailer) Send() error {
	auth := smtp.PlainAuth("", mail.User, mail.Password, mail.Host)
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"

	msg := []byte("To: " + mail.To + "\r\nFrom: " + mail.User + ">\r\nSubject: " + mail.Title + "\r\n" + content_type + "\r\n\r\n" + mail.Content)
	send_to := strings.Split(mail.To, ";")
	err := smtp.SendMail(mail.Host, auth, mail.User, send_to, msg)
	return err
}
