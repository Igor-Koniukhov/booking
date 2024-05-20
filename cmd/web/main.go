package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Igor-Koniukhov/bookings/internal/config"
	"github.com/Igor-Koniukhov/bookings/internal/driver"
	"github.com/Igor-Koniukhov/bookings/internal/handlers"
	"github.com/Igor-Koniukhov/bookings/internal/helpers"
	"github.com/Igor-Koniukhov/bookings/internal/models"
	"github.com/Igor-Koniukhov/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog
	err := godotenv.Load()
	if err == nil {
		log.Printf("Env load error: %v\n", err)
	}
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MailChan)
	listenForMail()
	fmt.Println("Starting mail listener...")

	fmt.Println(fmt.Sprintf("Starting application on port %s", os.Getenv("PORT")))

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = false
	app.InfoLog = helpers.InfoLog
	app.ErrorLog = helpers.ErrorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	db, err := driver.ConnectSQL(os.Getenv("DSN"))
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
