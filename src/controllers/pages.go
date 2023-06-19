// Package controllers provides ...
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/diegoparra/pet-adocao-frontend/src/answers"
	"github.com/diegoparra/pet-adocao-frontend/src/config"
	"github.com/diegoparra/pet-adocao-frontend/src/cookies"
	"github.com/diegoparra/pet-adocao-frontend/src/models"
	"github.com/diegoparra/pet-adocao-frontend/src/requests"
	"github.com/diegoparra/pet-adocao-frontend/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	// cookie, _ := cookies.Read(r)
	//
	// if cookie["token"] != "" {
	// 	http.Redirect(w, r, "/page/login", 302)
	// 	return
	// }

	utils.ExecTemplate(w, "login.html", nil)
}

func LoadRegisterUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] == "" {
		http.Redirect(w, r, "/page/login", 302)
		return
	}

	perfil := cookie["perfil"]

	if perfil != "Administrador" {
		answers.JSON(w, http.StatusForbidden, answers.Err{Err: "Nao Autorizado"})
		http.Redirect(w, r, "/admin/mostrar-usuarios", 302)
		return
	}

	utils.ExecTemplate(w, "cadastrar-usuarios.html", nil)
}

func LoadShowUsers(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] == "" {
		http.Redirect(w, r, "/page/login", 302)
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	fmt.Println(response)

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "mostrar-usuarios.html", users)
}

func LoadCreatePet(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "cadastrar-pet.html", nil)
}

func LoadHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/pet", config.APIURL)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var pet []models.Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "home.html", pet)
}

func LoadHomeEspecie(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	especie := parameters["especie"]
	url := fmt.Sprintf("%s/pet/%s", config.APIURL, especie)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var pet []models.Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "home.html", pet)
}

func LoadHomeEspecieAdmin(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	especie := parameters["especie"]
	url := fmt.Sprintf("%s/pet/%s", config.APIURL, especie)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var pet []models.Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "admin.html", pet)
}

func LoadGetPetById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	petID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	pet, err := models.SearchPetData(petID, r)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "pet-info.html", pet)
}

func LoadHomeAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] == "" {
		fmt.Println("Cookie not found")
		return
	}

	url := fmt.Sprintf("%s/pet", config.APIURL)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var pet []models.Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "admin.html", pet)
}

func LoadHomeAdotadosAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] == "" {
		fmt.Println("Cookie not found")
		return
	}

	url := fmt.Sprintf("%s/pet/adotados", config.APIURL)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		answers.TreatErrorStatusCode(w, response)
		return
	}

	var pet []models.Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "admin.html", pet)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchUserData(userID, r)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "profile.html", user)
}

func LoadEditPet(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	petID, _ := strconv.ParseUint(parameters["ID"], 10, 64)

	user, err := models.SearchPetData(petID, r)
	if err != nil {
		fmt.Println(err)
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "edit-pet.html", user)
}

func LoadUserEditProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] == "" {
		http.Redirect(w, r, "/page/login", 302)
		return
	}

	perfil := cookie["perfil"]

	if perfil != "Administrador" {
		answers.JSON(w, http.StatusForbidden, answers.Err{Err: "Nao Autorizado"})
		http.Redirect(w, r, "/admin/mostrar-usuarios", 302)
		return
	}

	parameters := mux.Vars(r)
	userID, _ := strconv.ParseUint(parameters["userID"], 10, 64)

	user, err := models.SearchUserData(userID, r)
	if err != nil {
		fmt.Println(err)
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	fmt.Println(user)

	utils.ExecTemplate(w, "editar-usuario.html", user)
}

func LoadChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "change-password.html", nil)
}

func LoadUserEditPhoto(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchUserData(userID, r)
	if err != nil {
		fmt.Println(err)
		answers.JSON(w, http.StatusInternalServerError, answers.Err{Err: err.Error()})
		return
	}

	utils.ExecTemplate(w, "edit-photo.html", user)
}
