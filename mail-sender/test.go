package main

import (
	"fmt"
	"net/smtp"
)

func NewMessage(topic, message string) []byte {
	str := fmt.Sprintf("Subject: %s\r\n", topic)
	str += fmt.Sprintf("\n%s\r\n", message)
	return []byte(str)
}

func test() {

	// Sender data.
	from := "bibadevelopment@gmail.com"
	password := "matb gcbx ivun kdyg"

	// Receiver email address.
	to := []string{
		"regffina.hisymutdinova@yandex.ru",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := "This is a test email message."
	msg := NewMessage("Topic", message)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}
