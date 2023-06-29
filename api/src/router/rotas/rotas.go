package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	funçao             func(http.ResponseWriter, *http.Request)
	requerAutenticação bool
}

// Configurar coloca todas as rotas no router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {

		if rota.requerAutenticação {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.funçao)),
				).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.funçao)).Methods(rota.Metodo)
		}
	}
	return r
}
