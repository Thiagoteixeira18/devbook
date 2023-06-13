package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da api
type Rota struct {
	URI                  string
	Metodo               string
	função               func(http.ResponseWriter, *http.Request)
	requerAutenticação   bool

	 
}
// Configurar coloa todas as rotas no router 
func Configurar (r *mux.Router) *mux.Router{
	rotas := rotasusuarios

	for _, rota:= range rotas{
		r.HandleFunc(rota.URI, rota.função).Methods(rota.Metodo)
	}
return r 
}