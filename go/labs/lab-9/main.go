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

	//
	problema2()

}

func problema2() {
	var size int = 10
	hashTable := hash_table.New[int](size, hash_table.HashFuncDivision)

	// Crear arreglo de numeros aleatorios
	var nums []int = crearArregloAleatorio(20)
	for _, num := range nums {
		println("Insertando numero: ", num)
		hashTable.Insert(num, num)
	}

	// Buscar un numero
	var num int = nums[13]
	println("Buscando numero: ", num)
	var result, _ = hashTable.Search(num)
	println("Resultado: ", result)

	// Eliminar un numero
	println("Eliminando numero: ", num)
	// hashTable.Delete(num)

}
