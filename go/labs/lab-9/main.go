package main

import (
	"data-structures/structures/hash_table"
)

func main() {

	println("\n\n--- Problema #2 ---\n\n")

	problema2()

	println("\n\n--- Problema #3 ---\n\n")

	problema3()
}

func problema2() {

	var size int = 10
	hashTable := hash_table.New[int](size, hash_table.HashFuncDivision)

	// Crear arreglo de numeros aleatorios
	var nums []int = crearArregloAleatorio(20)

	println("Insertando numeros: ")
	for _, num := range nums {
		print(num, ", ")
		hashTable.Insert(num, num)
	}

	println("\n")

	// Buscar un numero
	var num int = nums[13]
	println("Buscando numero... ", num)
	var result, _ = hashTable.Search(num)
	println("Resultado: ", result)

	// Eliminar un numero
	println("Eliminando numero: ", num)
	hashTable.Delete(num)

}

func problema3() {
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
