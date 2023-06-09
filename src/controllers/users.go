// Package controllers provides ...
package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/diegoparra/pet-adocao-frontend/src/answers"
	"github.com/diegoparra/pet-adocao-frontend/src/config"
	"github.com/diegoparra/pet-adocao-frontend/src/cookies"
	"github.com/diegoparra/pet-adocao-frontend/src/requests"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"nome":       r.FormValue("nome"),
		"email":      r.FormValue("email"),
		"nick":       r.FormValue("nick"),
		"senha":      r.FormValue("senha"),
		"nascimento": r.FormValue("nascimento"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
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

	answers.JSON(w, response.StatusCode, nil)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		return
	}

	cookies, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookies["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.APIURL, userID)

	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.TreatErrorStatusCode(w, response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"nome":       r.FormValue("nome"),
		"email":      r.FormValue("email"),
		"nick":       r.FormValue("nick"),
		"terapeuta":  r.FormValue("terapeuta"),
		"facebook":   r.FormValue("facebook"),
		"instragram": r.FormValue("instragram"),
		"telefone":   r.FormValue("telefone"),
	})

	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		fmt.Println(err)
		return
	}

	cookies, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookies["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.TreatErrorStatusCode(w, response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func EditUserPhoto(w http.ResponseWriter, r *http.Request) {

	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("assets/temp-images/", "upload-*.jpeg")
	if err != nil {
		fmt.Println("Erro ao criar tempFile")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao criar fileBytes")
		fmt.Println(err)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!

	fmt.Println(string(tempFile.Name()))

	user, err := json.Marshal(map[string]string{
		"file": string(tempFile.Name()),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		fmt.Println(err)
		return
	}

	cookies, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookies["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/photo", config.APIURL, userID)

	response, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.TreatErrorStatusCode(w, response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)

}
