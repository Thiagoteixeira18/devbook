package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		funçao:             controllers.CriarUsuario,
		requerAutenticação: false,
	},

	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		funçao:             controllers.BuscarUsuarios,
		requerAutenticação: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		funçao:             controllers.BuscarUsuario,
		requerAutenticação: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		funçao:             controllers.AtualizarUsuario,
		requerAutenticação: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		funçao:             controllers.DeltarUsuario,
		requerAutenticação: true,
	},
}
