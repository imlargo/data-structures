package main

import "math"

type Heap[T any] struct {
	data []int
	size int
}

func (heap *Heap[T]) getSize() int {
	return heap.size
}

func (heap *Heap[T]) parent(i int) int {
	return int(math.Ceil(float64(i)/2)) - 1
}

func (heap *Heap[T]) left(i int) int {
	return (2 * i) + 1
}

func (heap *Heap[T]) right(i int) int {
	return (2 * i) + 2
}

func (heap *Heap[T]) max_heapify(index int) []int {
	left := heap.left(index)
	right := heap.right(index)

	var largest int

	if left < heap.size && heap.data[left] > heap.data[index] {
		largest = left
	} else {
		largest = index
	}

	if right < heap.size && heap.data[right] > heap.data[largest] {
		largest = right
	}

	if largest != index {
		temp := heap.data[index]
		heap.data[index] = heap.data[largest]
		heap.data[largest] = temp
		heap.data = heap.max_heapify(largest)
	}

	return heap.data
}

func (heap *Heap[T]) build_max_heap() {

	var mid int = int(math.Floor(float64(heap.size) / 2))

	for i := mid; i >= 0; i-- {
		heap.data = heap.max_heapify(i)
	}
}

func (heap *Heap[T]) heap_sort() {
	heap.build_max_heap()

	for i := heap.size - 1; i >= 1; i-- {
		temp := heap.data[i]
		heap.data[i] = heap.data[0]
		heap.data[0] = temp

		heap.size = heap.size - 1
		heap.data = heap.max_heapify(0)
	}
}
