package main

import (
	"fmt"
	"net/smtp"
)

//sendmail
func SendMail(mail string, name string, surname string, subject string, cmdNumber string, message string) bool {

	to := ""

	//configure sending mailbox
	from := "service-recycleetvous@gmail.com"
	pass := "RessourceetVous2020"

	//set message
	message = name + " " + surname + "\n" + mail + "\n" + "Objet: " + subject + "\n" + "N commande: " + cmdNumber + "\n" + "Message: " + message

	msg := "From: " + mail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("smtp error: %s", err)
		fmt.Println("failed")
		return false
	} else {
		fmt.Println("passed")
		return true
	}
}
