// Package routes provides ...
package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Func:         controllers.DoLogin,
		AuthRequired: false,
	},
}
