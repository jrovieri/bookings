package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jrovieri/bookings/internal/config"
	"github.com/jrovieri/bookings/internal/db"
	"github.com/jrovieri/bookings/internal/handlers"
	"github.com/jrovieri/bookings/internal/helpers"
	"github.com/jrovieri/bookings/internal/models"
	"github.com/jrovieri/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	myDB, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer myDB.SQL.Close()

	defer close(app.MailChan)
	fmt.Println("Starting mail listener")

	listenForMail()

	// msg := models.MailData{
	// 	To:      "john.doe@example.com",
	// 	From:    "jrovieri@example.com",
	// 	Subject: "the subject",
	// 	Content: "",
	// }
	// app.MailChan <- msg

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Printf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*db.DB, error) {

	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	log.Println("Connectiong to database...")
	myDB, err := db.ConnectDB("host=localhost port=5432 dbname=bookings user=bookings password=gfq&?snAkpQx65p3")
	if err != nil {
		log.Fatal("cannot connect to database! stopping...")
		return nil, err
	}
	log.Println("connected to database!")

	tc, err := render.CreateTemplateCache()

	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorLog.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app, myDB)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return myDB, nil
}
