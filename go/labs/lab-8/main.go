package main

import (
	"data-structures/structures/tree"
)

func main() {
	var usuarios []*Usuario = make([]*Usuario, 6)
	usuarios[0] = &Usuario{Nombre: "Juan", Documento: "10101013"}
	usuarios[1] = &Usuario{Nombre: "Pablo", Documento: "10001011"}
	usuarios[2] = &Usuario{Nombre: "Maria", Documento: "10101015"}
	usuarios[3] = &Usuario{Nombre: "Ana", Documento: "1010000"}
	usuarios[4] = &Usuario{Nombre: "Diana", Documento: "10111105"}
	usuarios[5] = &Usuario{Nombre: "Mateo", Documento: "10110005"}

	arbolBinarioBusqueda := tree.BinarySearchTree[*Usuario]{}

	for _, usuario := range usuarios {
		println("Insertando usuario: ", usuario.GetKey())
		arbolBinarioBusqueda.Insert(usuario, usuario.GetKey())
	}

	arbolBinarioBusqueda.PrintInorder(arbolBinarioBusqueda.Root, func(u *Usuario) {
		println("Usuario: ", u.Nombre)
	})

	println(arbolBinarioBusqueda.Max(arbolBinarioBusqueda.Root).Nombre)
}
