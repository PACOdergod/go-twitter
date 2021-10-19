package main

import (
	"log"

	"github.com/PACOdergod/go-twitter/db"
	"github.com/PACOdergod/go-twitter/handlers"
)

func main() {
	if !db.ChequeoConnection() {
		log.Fatal("no hay conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
