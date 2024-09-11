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
		println("Insertando usuario: ", usuario.GetKey(), usuario.Nombre)
		arbolBinarioBusqueda.Insert(usuario, usuario.GetKey())
	}

	arbolBinarioBusqueda.PrintInorder(arbolBinarioBusqueda.Root, getNombreUsuario)

	// Metodos
	arbolBinarioBusqueda.PrintTree(getNombreUsuario)

	// Test Find
	var testKey int = usuarios[5].GetKey() // Mateo
	usuarioEncontrado := arbolBinarioBusqueda.Find(testKey)
	println("Usuario encontrado: ", usuarioEncontrado.Data.Value.Nombre)

	// Test Insert
	usuarioNuevo := &Usuario{Nombre: "Sebastian", Documento: "99999999"}
	arbolBinarioBusqueda.Insert(usuarioNuevo, usuarioNuevo.GetKey())
	arbolBinarioBusqueda.PrintTree(getNombreUsuario)

	// Test remove
	usuarioDiana := usuarios[4]
	println("Eliminando usuario: ", usuarioDiana.GetKey(), usuarioDiana.Nombre)
	arbolBinarioBusqueda.RemoveByKey(usuarioDiana.GetKey())
	arbolBinarioBusqueda.PrintTree(getNombreUsuario)

	arbolBinarioBusqueda.PrintInorder(arbolBinarioBusqueda.Root, getNombreUsuario)

}

func getNombreUsuario(u *Usuario) string {
	return u.Nombre
}
