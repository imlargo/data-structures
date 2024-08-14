package main

func main() {
	var heap Heap[int] = Heap[int]{data: []int{1, 2, 3, 4, 5}, size: 5}

	println("Heap size: ", heap.getSize())
}
