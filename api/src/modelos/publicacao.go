package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma piublicação de um usuario
type Publicacao struct {
	ID  uint64  `json:"id,omitempyt"`
	Titulo string `json:"titulo,omitempyt"`
	Conteudo string `json:"conteudo,omitempyt"`
	AutorID uint64 `json:"autorId,omitempyt"`
	AutorNick uint64  `json:"autorNick,omitempyt"`
	Curtidas uint64 `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm,omitempyt"`
}

// Preparar vai chamar os metodos para validar e formatar a publicação recebida
func (publicacao *Publicacao) Preparar() error{
	if erro := publicacao.validar(); erro != nil {
		return erro
	} 
	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o titulo é obrigatorio e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteudo é obrigatorio e não pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() error {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}