package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
