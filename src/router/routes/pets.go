package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routePets = []Route{
	{
		URI:          "/pet/cadastrar",
		Method:       http.MethodPost,
		Func:         controllers.CreatePet,
		AuthRequired: true,
	},
	{
		URI:          "/pet/cadastrar",
		Method:       http.MethodGet,
		Func:         controllers.LoadCreatePet,
		AuthRequired: true,
	},
	{
		URI:          "/pet/details/{ID}",
		Method:       http.MethodGet,
		Func:         controllers.LoadGetPetById,
		AuthRequired: false,
	},
	{
		URI:          "/pet/editar/{ID}",
		Method:       http.MethodGet,
		Func:         controllers.LoadEditPet,
		AuthRequired: false,
	},
	{
		URI:          "/pet/editar-animal/{ID}",
		Method:       http.MethodPut,
		Func:         controllers.UpdatePet,
		AuthRequired: false,
	},
}
