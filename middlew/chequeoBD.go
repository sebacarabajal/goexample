package middlew

import (
	"net/http"

	"github.com/sebacarabajal/goexample/bd"
)

//los middlewares reciben y envian el mismo tipo de objeto, es un pasamano
//este solo chequea la conexión con la base y pasa el control al siguiente
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	//devuelve función anónima
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida base de datos", 500)
			return
		}

		//Devuelve función donde paso al próximo eslabon todos los valores que he recibido
		next.ServeHTTP(w, r)
	}
}
