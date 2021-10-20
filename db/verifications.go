package db

import (
	"context"
	"time"

	"github.com/PACOdergod/go-twitter/models"

	"go.mongodb.org/mongo-driver/bson"
)

// TODO: terminar
func ExisteUsuario(email string) (bool, models.Usuario, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, resultado, ""
	}

	ID := resultado.ID.Hex()
	return true, resultado, ID
}
