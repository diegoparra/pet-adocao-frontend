package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"webapp/src/answers"
	"webapp/src/config"
	"webapp/src/requests"

	"github.com/gorilla/mux"
)

func CreatePet(w http.ResponseWriter, r *http.Request) {

	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("arquivo")
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
	tempFile, err := ioutil.TempFile("assets/temp-images", "upload-*.jpeg")
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

	pet, err := json.Marshal(map[string]string{
		"nome":      r.FormValue("nome"),
		"telefone":  r.FormValue("telefone"),
		"especie":   r.FormValue("especie"),
		"genero":    r.FormValue("genero"),
		"porte":     r.FormValue("porte"),
		"vacinado":  r.FormValue("vacinado"),
		"castrado":  r.FormValue("castrado"),
		"descricao": r.FormValue("descricao"),
		"arquivo":   string(tempFile.Name()),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/pet/cadastrar", config.APIURL)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(pet))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error doing the request")
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

func UpdatePet(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	petID, err := strconv.ParseUint(parameters["ID"], 10, 64)

	pet, err := json.Marshal(map[string]string{
		"nome":      r.FormValue("nome"),
		"telefone":  r.FormValue("telefone"),
		"especie":   r.FormValue("especie"),
		"genero":    r.FormValue("genero"),
		"porte":     r.FormValue("porte"),
		"vacinado":  r.FormValue("vacinado"),
		"castrado":  r.FormValue("castrado"),
		"descricao": r.FormValue("descricao"),
		"adotado":   r.FormValue("adotado"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/pet/editar/%d", config.APIURL, petID)
	response, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(pet))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error doing the request")
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
