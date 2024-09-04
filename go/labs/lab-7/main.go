package main

import (
	"data-structures/structures/heap"
	"math/rand/v2"
)

func main() {

	println("#  - - - Pruebas de Heap - - -  #\n")
	pruebaHeap()

	println("\n\n\n#  - - - Pruebas de PriorityQueue - - -  #\n")
	pruebaPriorityQueue()

}

func pruebaHeap() {
	// Implementacion heap
	var cantidad int = 20

	var data []int = crearArregloAleatorio(cantidad)

	// Inicializar el heap
	println("Heap inicializado aleatoriamente: ")
	heap := heap.Heap{Data: data, Size: cantidad}
	heap.Print()

	// Construir el heap
	println("\nBuild Max Heap: ")
	heap.Build_max_heap()
	heap.Print()

	// Heap Sort
	println("\nHeap Sort: ")
	heap.Heap_sort()
	heap.Print()
}

func pruebaPriorityQueue() {
	// Implementacion heap
	var cantidad int = 20

	// Crear arreglo de numeros aleatorios
	var data []int = crearArregloAleatorio(cantidad)

	// Inicializar el heap
	println("Heap inicializado aleatoriamente: ")
	priotiryQueue := heap.Heap{Data: data, Size: cantidad}
	priotiryQueue.Print()

	// Construir el heap
	println("\nBuild Max Heap: ")
	priotiryQueue.Build_max_heap()
	priotiryQueue.Print()

	// Max Heap Insert
	println("\nMax Heap Insert: ", 101)
	priotiryQueue.MaxHeapInsert(101)
	priotiryQueue.Print()

	// Heap Extract Max
	println("\nHeap Extract Max: ")
	max := priotiryQueue.HeapExtractMax()
	println("Maximo extraido: ", max)
	priotiryQueue.Print()

	// Heap Maximum
	println("\nHeap Maximum: ")
	max = priotiryQueue.HeapMaximum()
	println("Maximo: ", max)

}

func crearArregloAleatorio(cantidad int) []int {
	// Crear arreglo de numeros aleatorios
	var data []int = make([]int, cantidad)
	for i := 0; i < cantidad; i++ {
		var num int = rand.IntN(100)
		data[i] = num
	}

	return data
}
