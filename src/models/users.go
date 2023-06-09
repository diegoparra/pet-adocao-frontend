package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"nome"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	Birth     string    `json:"nascimento"`
	Terapeuta string    `json:"terapeuta"`
	Facebook  string    `json:"facebook"`
	Instagram string    `json:"instagram"`
	Telefone  string    `json:"telefone"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"criadoEm"`
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
