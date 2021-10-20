package db

import (
	"context"
	"time"

	"github.com/PACOdergod/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// inserta registro en la base de datos
func CreateUser(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// cuando termine la ejecucion de la funcion cancelara el timeout
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = Encriptar(u.Password)
	// TODO: que pasa cuando hubo un error al encriptar

	// inserta el usuario en la coleccion
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	// para obtener el id del registro
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
