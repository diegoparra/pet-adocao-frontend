package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routePages = []Route{
	{
		URI:          "/page/login",
		Method:       http.MethodGet,
		Func:         controllers.LoadLoginPage,
		AuthRequired: false,
	},
}
