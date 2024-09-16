package main

import (
	"data-structures/structures/hash_table"
)

func main() {

	var size int = 5
	var usuarios []*Usuario = getUsuarios()

	/* Crear tabla hash con HashFuncDivision */
	hashTable1 := hash_table.New[*Usuario](size, hash_table.HashFuncDivision)
	for _, usuario := range usuarios {
		println("Insertando usuario: ", usuario.Id, usuario.Nombre)
		hashTable1.Insert(usuario.Id, usuario)
	}
	// Mostrar contadores
	for i, lista := range hashTable1.Array {
		println(i, " -> ", lista.Size)
	}

	/* Crear tabla hash con HashFuncMultiplicacion */
	hashTable2 := hash_table.New[*Usuario](size, hash_table.HashFuncMultiplicacion)
	for _, usuario := range usuarios {
		println("Insertando usuario: ", usuario.Id, usuario.Nombre)
		hashTable2.Insert(usuario.Id, usuario)
	}
	// Mostrar contadores
	for i, lista := range hashTable2.Array {
		println(i, " -> ", lista.Size)
	}

}
