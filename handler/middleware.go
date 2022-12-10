package handler

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	})
}

func basicAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path == "/api/v1/user/login" {
			next(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error": "Unauthorized"}`)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		next(w, r)
	})
}
