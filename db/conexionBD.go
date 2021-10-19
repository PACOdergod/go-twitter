package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uriDB = "mongodb+srv://twitero:holamundo@cluster0.w13mm.mongodb.net/Cluster0?retryWrites=true&w=majority"
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI(uriDB)

// ConectarDB permite conectar a la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa")
	return client
}

// ChequeoConnection verifica que la base de datos esta funcionando
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
