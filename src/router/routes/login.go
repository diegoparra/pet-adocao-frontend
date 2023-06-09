// Package routes provides ...
package routes

import (
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/controllers"
)

var routesLogin = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Func:         controllers.DoLogin,
		AuthRequired: false,
	},
}
