package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Indirect
)

// Conectar abre a conexção com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringDeConexãoBanco)
	if erro != nil {
		return nil, erro
	}
	return db, nil
}
