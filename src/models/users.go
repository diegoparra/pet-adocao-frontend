package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diegoparra/pet-adocao-frontend/src/config"
	"github.com/diegoparra/pet-adocao-frontend/src/requests"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	Perfil   string    `json:"perfil,omitempty"`
	Ativo    string    `json:"ativo,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func SearchUserData(userID uint64, r *http.Request) (User, error) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		fmt.Println(err)
		return User{}, err
	}

	return user, nil
}
