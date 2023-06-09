package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/answers"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
)

// DoLogin use the email and password to authenticate a user into the app
func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var dataAuthentication models.DataAuthentication

	if err := json.NewDecoder(response.Body).Decode(&dataAuthentication); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, dataAuthentication.ID, dataAuthentication.Token); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

}
