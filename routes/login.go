package routes

import (
	"encoding/json"
	"net/http"

	"github.com/PACOdergod/go-twitter/db"
	"github.com/PACOdergod/go-twitter/jwt"
	"github.com/PACOdergod/go-twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario o contraseña invalida"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña invalida", 400)
		return
	}

	existe, user := db.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario o contraseña invalidos", 400)
		return
	}

	// como logro logear el usuario devolvera un token
	// JWT
	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el JWT"+err.Error(), 400)
	}

	resp := models.RespLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// como grabar una cookie
	// expirationTime := time.Now().Add(24 * time.Hour)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   jwtKey,
	// 	Expires: expirationTime,
	// })
}
