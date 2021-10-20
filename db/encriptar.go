package db

import "golang.org/x/crypto/bcrypt"

func Encriptar(s string) (string, error) {

	costo := 7
	// primero debe convertir el string a bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), costo)
	return string(bytes), err
}
