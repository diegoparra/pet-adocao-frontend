package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diegoparra/pet-adocao-frontend/src/config"
	"github.com/diegoparra/pet-adocao-frontend/src/requests"
)

type Pet struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Telefone  string    `json:"telefone,omitempty"`
	Especie   string    `json:"especie,omitempty"`
	Descricao string    `json:"descricao,omitempty"`
	Genero    string    `json:"genero,omitempty"`
	Porte     string    `json:"porte,omitempty"`
	Vacinado  string    `json:"vacinado,omitempty"`
	Castrado  string    `json:"castrado,omitempty"`
	Adotado   string    `json:"adotado,omitempty"`
	Arquivo   string    `json:"arquivo,omitempty"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

func SearchPetData(petID uint64, r *http.Request) (Pet, error) {

	url := fmt.Sprintf("%s/pet/details/%d", config.APIURL, petID)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao fazer o request")
		return Pet{}, err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Erro no response status code")
		return Pet{}, err
	}

	var pet Pet
	if err = json.NewDecoder(response.Body).Decode(&pet); err != nil {
		fmt.Println("erro ao fazer o Decode")
		return Pet{}, err
	}

	return pet, nil
}
