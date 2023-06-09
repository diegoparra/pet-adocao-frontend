// Package routes provides ...
package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeUsers = []Route{
	{
		URI:          "/create-user",
		Method:       http.MethodGet,
		Func:         controllers.LoadRegisterUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserProfile,
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
		URI:          "/edit-profile",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserEditProfile,
		AuthRequired: true,
	},
	{
		URI:          "/edit-profile",
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
