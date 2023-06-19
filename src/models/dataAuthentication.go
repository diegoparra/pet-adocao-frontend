package models

// DataAuthentication has the id and token from the authenticated user
type DataAuthentication struct {
	ID     string `json:"id"`
	Token  string `json:"token"`
	Perfil string `json:"perfil"`
}
