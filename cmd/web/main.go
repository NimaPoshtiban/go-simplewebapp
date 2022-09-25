package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//! Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false //! change this if you ever wanted to host this app

	app.Session = session

	tc, err := render.CreateTemplateCache()

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	fmt.Printf("Starting app on port %s\n", portNumber[1:])

	srv := &http.Server{
		Addr:        portNumber,
		Handler:     routes(&app),
		IdleTimeout: time.Second * 5,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
