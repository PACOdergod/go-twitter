package jwt

import (
	"time"

	"github.com/PACOdergod/go-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.Usuario) (string, error) {
	miClave := []byte("MasterGO")

	//! nunca guardar la password en un jwt
	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"apellidos":        user.Apellidos,
		"fecha_nacimiento": user.Nacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
