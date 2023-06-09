package routes

import (
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/controllers"
)

var routeLogout = Route{
	URI:          "/users/sair",
	Method:       http.MethodGet,
	Func:         controllers.Logout,
	AuthRequired: true,
}
