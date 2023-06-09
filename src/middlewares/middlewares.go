package middlewares

import (
	"log"
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Read(r)
		if err != nil {
			http.Redirect(w, r, "/page/login", 302)
			return
		}
		next(w, r)
	}
}
