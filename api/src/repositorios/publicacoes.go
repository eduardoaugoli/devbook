package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar inseri uma publicacao no banco
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo,conteudo,autor_id) values (?,?,?)",
	)
	if err != nil {
		return 0, err
	}

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, err := repositorio.db.Query(
		"SELECT p.id,p.titulo,p.conteudo,p.autor_id,p.curtidas,p.criadaEm,u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?", publicacaoID,
	)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}
	return publicacao, nil

}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		"SELECT DISTINCT p.id,p.titulo,p.conteudo,p.autor_id,p.curtidas,p.criadaEm,u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id inner join seguidores s on s.usuario_id = p.autor_id where u.id = ? or s.seguidor_id = ?", usuarioID, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}
