package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)
//That is a common format of most middleware. All of it writing in a similar way

// NoSurf adds CSRF protection to all POST requests
// CSRFToken - cross site request forgery token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler:= nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,

	})
	return csrfHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler  {
	return session.LoadAndSave(next)

}