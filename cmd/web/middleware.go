package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// ? Dummy middleware for showing how middleware can be implemented
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf Generates CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, //! Change this to true when in Production
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad saves and loads session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
