package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting cookies")
	cookies.Del(w)
	http.Redirect(w, r, "/", 302)
}
