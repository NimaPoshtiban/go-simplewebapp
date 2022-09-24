package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

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
