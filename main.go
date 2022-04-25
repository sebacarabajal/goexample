package main

import (
	"log"

	"github.com/sebacarabajal/goexample/bd"
	"github.com/sebacarabajal/goexample/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}
