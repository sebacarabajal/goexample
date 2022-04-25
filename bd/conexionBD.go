package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://zeeeee:Seba5854!@cluster0.5uwxr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD() Para conectar la base de datos*/
func ConectarBD() *mongo.Client {
	//TODO : Contexto. Conecte a la base de datos sin ningún timeout ni nada
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
	log.Println("Conexión exitosa")

	return client

}

/*ChequeoConnection()*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
