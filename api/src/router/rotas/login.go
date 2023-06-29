package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:  				"/login",
	Metodo:  			http.MethodPost,
	funçao:  			controllers.Login,
	requerAutenticação: false,
}