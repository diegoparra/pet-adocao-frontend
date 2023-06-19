// Package routes provides ...
package routes

import (
	"net/http"

	"github.com/diegoparra/pet-adocao-frontend/src/controllers"
)

var routeUsers = []Route{
	{
		URI:          "/admin/cadastrar-usuario",
		Method:       http.MethodGet,
		Func:         controllers.LoadRegisterUser,
		AuthRequired: true,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/admin/mostrar-usuarios",
		Method:       http.MethodGet,
		Func:         controllers.LoadShowUsers,
		AuthRequired: true,
	},
	{
		URI:          "/change-password",
		Method:       http.MethodGet,
		Func:         controllers.LoadChangePasswordPage,
		AuthRequired: true,
	},
	{
		URI:          "/change-password",
		Method:       http.MethodPost,
		Func:         controllers.ChangePassword,
		AuthRequired: true,
	},
	{
		URI:          "/page/editar-usuario/{userID}",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserEditProfile,
		AuthRequired: true,
	},
	{
		URI:          "/editar-usuario/{userID}",
		Method:       http.MethodPut,
		Func:         controllers.EditUserProfile,
		AuthRequired: true,
	},
	{
		URI:          "/edit-photo",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserEditPhoto,
		AuthRequired: true,
	},
	{
		URI:          "/edit-photo",
		Method:       http.MethodPut,
		Func:         controllers.EditUserPhoto,
		AuthRequired: true,
	},
}
