package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// cria token com permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	premissoes := jwt.MapClaims{}

	premissoes["authorized"] = true
	premissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	premissoes["usuarioID"] = usuarioID
	//secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, premissoes)
	return token.SignedString([]byte(config.SecretKey)) //secret
}
