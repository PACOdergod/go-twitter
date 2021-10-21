package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/PACOdergod/go-twitter/middleware"
	"github.com/PACOdergod/go-twitter/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	// ENDPOINTS
	// ruta para registro de usuarios
	// cuando en el navegador se busque /registros
	// se mandara a llamar la funcion que se le pase
	// pero la peticion debe ser POST
	router.HandleFunc("/registro", middleware.ChequeoDB(routes.Registro)).Methods("POST")

	router.HandleFunc("/login", middleware.ChequeoDB(routes.Login)).Methods("POST")

	// sirve para saber si la maquina tiene ya seteado
	// algun puerto, si no lo tiene lo seteo
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// cors sirve para permitir peticiones remotas
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
