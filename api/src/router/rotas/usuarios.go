package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasusuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		função:             controllers.CriarUsuario,
		requerAutenticação: false,
	},

{
	URI:    "/usuarios",
	Metodo: http.MethodGet,
	função: controllers.BuscarUsuarios,
	requerAutenticação: false,
},
	{
	URI:    "/usuarios/{usuarioId}",
	Metodo: http.MethodGet,
	função: controllers.BuscarUsuario,
	requerAutenticação: false,
},
{
	URI:    "/usuarios/{usuarioId}",
	Metodo: http.MethodPut,
	função: controllers.AtualizarUsuario,
	requerAutenticação: false,
},
	{
	URI:    "/usuarios/{usuarioId}",
	Metodo: http.MethodDelete,
	função: controllers.DeltarUsuario,
	requerAutenticação: false,
},
}