package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeLogout = Route{
	URI:          "/users/sair",
	Method:       http.MethodGet,
	Func:         controllers.Logout,
	AuthRequired: true,
}
