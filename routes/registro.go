package routes

import (
	"encoding/json"
	"net/http"

	"github.com/PACOdergod/go-twitter/db"
	"github.com/PACOdergod/go-twitter/models"
)

// creara un usuraio en la base de datos
// el r.body es un objeto de lectura unica
// cuando se lee se detruye
func Registro(w http.ResponseWriter, r *http.Request) {

	// decodificara el json con los datos que venga en la peticion
	// y los asignara a una variable tipo usuario
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error el los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser minimo 6 caracteres", 400)
		return
	}

	existeUsuario, _, _ := db.ExisteUsuario(t.Email)
	if existeUsuario {
		http.Error(w, "Ya esta registrado el email", 400)
		return
	}

	_, status, err := db.CreateUser(t)
	if err != nil {
		http.Error(w, "Error al intentar registrar usuario"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se pudo registrar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
