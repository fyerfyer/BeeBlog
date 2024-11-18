package models

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Email struct {
	From    string
	To      []string
	Subject string
	HTML    []byte
}

var emailSender = email.NewEmail()

func SendEmail(email Email) error {
	emailSender.From = email.From
	emailSender.To = email.To
	emailSender.Subject = email.Subject
	emailSender.HTML = email.HTML

	return emailSender.Send("smtp.gmail.com:587",
		smtp.PlainAuth("", "fuy60703@gmail.com", "jyqcbigrajyrrcnn", "smtp.gmail.com"))
}
