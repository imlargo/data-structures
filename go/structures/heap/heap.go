package heap

import "math"

type Heap[T any] struct {
	Data []int
	Size int
}

func (heap *Heap[T]) GetSize() int {
	return heap.Size
}

func (heap *Heap[T]) Parent(i int) int {
	return int(math.Ceil(float64(i)/2)) - 1
}

func (heap *Heap[T]) Left(i int) int {
	return (2 * i) + 1
}

func (heap *Heap[T]) Right(i int) int {
	return (2 * i) + 2
}

func (heap *Heap[T]) Max_heapify(index int) []int {
	left := heap.Left(index)
	right := heap.Right(index)

	var largest int

	if left < heap.Size && heap.Data[left] > heap.Data[index] {
		largest = left
	} else {
		largest = index
	}

	if right < heap.Size && heap.Data[right] > heap.Data[largest] {
		largest = right
	}

	if largest != index {
		temp := heap.Data[index]
		heap.Data[index] = heap.Data[largest]
		heap.Data[largest] = temp
		heap.Data = heap.Max_heapify(largest)
	}

	return heap.Data
}

func (heap *Heap[T]) Build_max_heap() {

	var mid int = int(math.Floor(float64(heap.Size) / 2))

	for i := mid; i >= 0; i-- {
		heap.Data = heap.Max_heapify(i)
	}
}

func (heap *Heap[T]) Heap_sort() {
	heap.Build_max_heap()

	for i := heap.Size - 1; i >= 1; i-- {
		temp := heap.Data[i]
		heap.Data[i] = heap.Data[0]
		heap.Data[0] = temp

		heap.Size = heap.Size - 1
		heap.Data = heap.Max_heapify(0)
	}
}
