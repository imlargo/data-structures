package main

import "math"

type Heap[T any] struct {
	data []T
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

func (heap *Heap[T]) max_heapify(array []int, index int, size int) {
	println("Max heapify")
}
func (heap *Heap[T]) build_max_heap() {
	println("Building max heap")
}
func (heap *Heap[T]) heap_sort() {
	println("Heap sort")
}
