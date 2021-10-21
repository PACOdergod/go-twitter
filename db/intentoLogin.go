package db

import (
	"github.com/PACOdergod/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (bool, models.Usuario) {
	existeU, usu, _ := ExisteUsuario(email)
	if !existeU {
		return false, usu
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return false, usu
	}

	return true, usu
}
