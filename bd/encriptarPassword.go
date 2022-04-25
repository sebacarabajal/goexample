package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {

	//cantidad de pasadas sobre el texto para encriptarlo. Más tiempo
	costo := 8

	//solo acepta slices
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err

}
