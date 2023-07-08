package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		funçao:             controllers.CriarPublicacao,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		funçao:             controllers.BuscarPublicacoes,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		funçao:             controllers.AtualizarPublicacao,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		funçao:				controllers.AtualizarPublicacao,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		funçao:             controllers.DeletarPublicacao,
		requerAutenticação: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		funçao:             controllers.BuscarPublicacoesPorUsuario,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		funçao:             controllers.CurtirPublicacao,
		requerAutenticação: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		funçao:             controllers.DescurtirPublicacao,
		requerAutenticação: true,
	},
}
