package modelos

// Contem o id e o token do usuario
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
