package main

import (
	"data-structures/structures/hash_table"
)

func main() {

	println("\n\n# - - - Problema #2 - - - #\n\n")

	println("\n--- Usando hash de division ---\n")
	problema2(hash_table.HashFuncDivision)

	println("\n--- Usando hash de multiplicacion ...\n")
	problema2(hash_table.HashFuncMultiplicacion)

	println("\n\n--- Problema #3 ---\n\n")

	println("\n--- Usando hash de division ---\n")
	problema3(hash_table.HashFuncDivision)

	println("\n--- Usando hash de multiplicacion ...\n")
	problema3(hash_table.HashFuncMultiplicacion)

}

func problema2(selectedHashFunc func(int, int) int) {

	var size int = 10
	hashTable := hash_table.New[int](size, selectedHashFunc)

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
	println("Eliminando numero", num, "en index", hashTable.HashFunc(num, hashTable.Size))
	hashTable.Delete(num)

	// Mostrar distribucion
	hashTable.MostrarDistribucion()

}

func problema3(selectedHashFunc func(int, int) int) {
	var size int = 5
	var usuarios []*Usuario = getUsuarios()

	/* Crear tabla hash con HashFuncDivision */
	hashTable := hash_table.New[*Usuario](size, selectedHashFunc)
	for _, usuario := range usuarios {
		println("Insertando usuario: ", usuario.Id, usuario.Nombre)
		hashTable.Insert(usuario.Id, usuario)
	}
	// Mostrar distribucion
	hashTable.MostrarDistribucion()
}
