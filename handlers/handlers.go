package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors" //permisos api
	"github.com/sebacarabajal/goexample/middlew"
	"github.com/sebacarabajal/goexample/routers"
)

//Con los punteros no modificamos el valor de la variable, sino el valor de ese espacio de memoria
//Con el * accedemos al valor de ese espacio de memoria
/*
var creature string = "shark"
var pointer *string = &creature

fmt.Println("creature =", creature)
fmt.Println("pointer =", pointer)
fmt.Println("*pointer =", *pointer)
*/

func Manejadores() {
	//mux captura el http y procesa los request y response
	router := mux.NewRouter()

	//definimos endpoints
	//Cuando detecte que hay un request a /registro usando POST, llama a un middlew, que
	//chquea la conexión y si todo está ok devuelve el control a routers
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	//Vemos si en el OS ya hay una variable de entorno. Prioridad
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	//Si tengo permisos me deja continuar
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
