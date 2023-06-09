package routes

import (
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/controllers"
)

var routeHome = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Func:         controllers.LoadHome,
		AuthRequired: false,
	},
	{
		URI:          "/{especie}",
		Method:       http.MethodGet,
		Func:         controllers.LoadHomeEspecie,
		AuthRequired: false,
	},
	{
		URI:          "/users/admin",
		Method:       http.MethodGet,
		Func:         controllers.LoadHomeAdmin,
		AuthRequired: true,
	},
	{
		URI:          "/admin/{especie}",
		Method:       http.MethodGet,
		Func:         controllers.LoadHomeEspecieAdmin,
		AuthRequired: true,
	},
	{
		URI:          "/admin/adotados",
		Method:       http.MethodGet,
		Func:         controllers.LoadHomeAdotadosAdmin,
		AuthRequired: true,
	},
}
