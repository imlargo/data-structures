package main

import "strconv"

type Usuario struct {
	Nombre    string
	Documento string
}

func (usuario *Usuario) GetKey() int {

	cadena := usuario.Documento

	var key int = 0
	for _, v := range cadena {
		num, _ := strconv.Atoi(string(v))
		key += num
	}

	return key

}
