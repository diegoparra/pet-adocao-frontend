package routes

import (
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/controllers"
)

var routePages = []Route{
	{
		URI:          "/page/login",
		Method:       http.MethodGet,
		Func:         controllers.LoadLoginPage,
		AuthRequired: false,
	},
	{
		URI:          "/page/cadastrar-usuario",
		Method:       http.MethodGet,
		Func:         controllers.LoadRegisterUser,
		AuthRequired: false,
	},
}
