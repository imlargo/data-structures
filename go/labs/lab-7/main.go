package main

import (
	"data-structures/structures/heap"
	"math/rand/v2"
)

func main() {

	println("Pruebas de Heap\n")
	pruebaHeap()

	println("\n\nPruebas de PriorityQueue\n")
	pruebaPriorityQueue()

}

func pruebaHeap() {
	// Implementacion heap
	var cantidad int = 20

	// Crear arreglo de numeros aleatorios
	var data []int = make([]int, cantidad)
	for i := 0; i < cantidad; i++ {
		var num int = rand.IntN(100)
		data[i] = num
	}

	// Inicializar el heap
	println("Heap inicializado: ")
	heap := heap.Heap{Data: data, Size: cantidad}
	heap.Print()

	// Construir el heap
	println("Build Max Heap: ")
	heap.Build_max_heap()
	heap.Print()

	// Heap Sort
	println("Heap Sort: ")
	heap.Heap_sort()
	heap.Print()
}

func pruebaPriorityQueue() {
	// Implementacion heap
	var cantidad int = 20

	// Crear arreglo de numeros aleatorios
	var data []int = make([]int, cantidad)
	for i := 0; i < cantidad; i++ {
		var num int = rand.IntN(100)
		data[i] = num
	}

	// Inicializar el heap
	println("Heap inicializado: ")
	priotiryQueue := heap.Heap{Data: data, Size: cantidad}
	priotiryQueue.Print()

	// Construir el heap
	println("Build Max Heap: ")
	priotiryQueue.Build_max_heap()
	priotiryQueue.Print()

	// Heap Sort
	println("Heap Sort: ")
	priotiryQueue.Heap_sort()
	priotiryQueue.Print()

	// Max Heap Insert
	println("Max Heap Insert: ")
	priotiryQueue.MaxHeapInsert(100)
	priotiryQueue.Print()

	// Heap Extract Max
	println("Heap Extract Max: ")
	max := priotiryQueue.HeapExtractMax()
	println("Maximo extraido: ", max)
	priotiryQueue.Print()

	// Heap Maximum
	println("Heap Maximum: ")
	max = priotiryQueue.HeapMaximum()
	println("Maximo: ", max)

}
