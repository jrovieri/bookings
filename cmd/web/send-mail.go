package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jrovieri/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {

	go func() {
		for {
			msg := <-app.MailChan
			sendMessage(msg)
		}
	}()
}

func sendMessage(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, string(m.Content))
	} else {
		data, err := os.ReadFile(fmt.Sprintf("./templates/mail/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		message := strings.Replace(string(data), "[%body%]", string(m.Content), 1)
		email.SetBody(mail.TextHTML, message)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email sent!")
	}
}
