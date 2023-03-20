package modelos

// Senha representa o formato da requisicao de alteracao de senha
type Senha struct {
	Nova  string `json:nova`
	Atual string `json:atual`
}
