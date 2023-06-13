package controllers

import (
	
	"net/http"
)

// CriarUsuario cria um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuario"))
}

// BucandoUsuarios buca um usuario no banco de dados 
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando Todos os Usuario"))
}

// BuscarUsuario busca um usuario no banco de dados 
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando Usuario"))
}

// AtualizarUsuario atualiza um usuario no banco de dados 
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuario"))
}

// DeletarUsuario deleta um usuario no banco de dados 
func DeltarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario"))
}