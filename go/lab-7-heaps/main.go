package main

import (
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

	var heap Heap[int] = Heap[int]{data: data, size: cantidad}

	heap.build_max_heap()

	// Imprimir el arreglo
	println("Max heap: ")
	for i := 0; i < heap.size; i++ {
		print(heap.data[i], ", ")
	}

	println("\n\n")

	println("Heap size: ", heap.getSize())
}
