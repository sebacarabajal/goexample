package bd

import (
	"context"
	"time"

	"github.com/sebacarabajal/goexample/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Mayuscula asi la puedo exportar

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	//formateo a formato bson
	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
