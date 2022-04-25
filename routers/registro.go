package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebacarabajal/goexample/bd"
	"github.com/sebacarabajal/goexample/models"
)

//recibe el json con los datos del usuario que se quiere registrar
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email obligatorio ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Contraseña demasiado corta ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe el usuario ", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Error al insertar registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el usuario ", 400)
		return
	}

	//escribo en el header
	w.WriteHeader(http.StatusCreated)

}
