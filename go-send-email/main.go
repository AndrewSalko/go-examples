package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	//Based on https://mailtrap.io/blog/golang-send-email/

	fmt.Println("Go send email SMTP sample app")

	emailAddr := "youremail@ukr.net"
	emailPass := "<your pass>"
	smtpHost := "smtp.ukr.net"

	message := gomail.NewMessage()

	message.SetHeader("From", emailAddr)
	message.SetHeader("To", "anothermail@gmail.com")
	message.SetHeader("Subject", "Hello from Andrew Gomail")

	message.SetBody("text/plain", "Hello! Go is great!")

	d := gomail.NewDialer(smtpHost, 465, emailAddr, emailPass)

	//if SSL/TLS certificate is not valid on server (in prod.set to false)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//Send
	err := d.DialAndSend(message)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Finished")
} //main
