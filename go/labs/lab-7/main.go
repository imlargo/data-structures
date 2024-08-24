package main

import (
	"data-structures/structures/heap"
	"math/rand/v2"
)

func main() {

	var cantidad int = 20

	var data []int = make([]int, cantidad)
	for i := 0; i < cantidad; i++ {
		var num int = rand.IntN(100)
		data[i] = num
	}

	// Imprimir el arreglo
	println("Arreglo: ")
	for i := 0; i < cantidad; i++ {
		print(data[i], ", ")
	}

	println("\n\n")

	heap := heap.Heap{Data: data, Size: cantidad}

	heap.Build_max_heap()

	// Imprimir el arreglo
	println("Max heap: ")
	for i := 0; i < heap.Size; i++ {
		print(heap.Data[i], ", ")
	}

	println("\n\n")

	heap.Heap_sort()

	println("Heap size: ", heap.GetSize())

	// Imprimir el arreglo
	println("Heapify: ")
	for i := 0; i < heap.Size; i++ {
		print(heap.Data[i], ", ")
	}

}
