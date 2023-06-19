package cookies

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"

	"github.com/diegoparra/pet-adocao-frontend/src/config"
)

var s *securecookie.SecureCookie

// Config use the env var to create the secure cookie
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save register auth information
func Save(w http.ResponseWriter, ID, token string, perfil string) error {
	data := map[string]string{
		"id":     ID,
		"token":  token,
		"perfil": perfil,
	}

	dataCoded, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataCoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func Del(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
