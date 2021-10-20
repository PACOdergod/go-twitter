package middleware

import (
	"net/http"

	"github.com/PACOdergod/go-twitter/db"
)

// los middleware tiene que devolver lo mismo que reciben

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		// si hay un problema con la base de datos se detendra
		// y mostrara un error
		if !db.ChequeoConnection() {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}

		// si no, dejara pasar toda la informacion
		next.ServeHTTP(rw, r)
	}
}
