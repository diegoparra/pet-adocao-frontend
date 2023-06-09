package controllers

import (
	"fmt"
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting cookies")
	cookies.Del(w)
	http.Redirect(w, r, "/", 302)
}
