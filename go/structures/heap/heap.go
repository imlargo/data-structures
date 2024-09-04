package heap

import (
	"math"
	"strconv"
)

type Heap struct {
	Data []int
	Size int
}

func (heap *Heap) GetSize() int {
	return heap.Size
}

func (heap *Heap) Parent(i int) int {
	return int(math.Ceil(float64(i)/2)) - 1
}

func (heap *Heap) Left(i int) int {
	return (2 * i) + 1
}

func (heap *Heap) Right(i int) int {
	return (2 * i) + 2
}

func (heap *Heap) Max_heapify(index int) []int {
	left := heap.Left(index)
	right := heap.Right(index)

	var largest int = index

	if left < heap.Size && heap.Data[left] > heap.Data[index] {
		largest = left
	}

	if right < heap.Size && heap.Data[right] > heap.Data[largest] {
		largest = right
	}

	if largest != index {
		// Intercambiar los valores
		heap.Data[index], heap.Data[largest] = heap.Data[largest], heap.Data[index]

		heap.Data = heap.Max_heapify(largest)
	}

	return heap.Data
}

func (heap *Heap) Build_max_heap() {

	var mid int = int(math.Floor(float64(heap.Size) / 2))

	for i := mid; i >= 0; i-- {
		heap.Data = heap.Max_heapify(i)
	}
}

func (heap *Heap) Heap_sort() {
	heap.Build_max_heap()

	for i := heap.Size - 1; i >= 1; i-- {
		// Intercambiar los valores
		heap.Data[i], heap.Data[0] = heap.Data[0], heap.Data[i]

		heap.Size -= 1
		heap.Data = heap.Max_heapify(0)
	}

	heap.Size = len(heap.Data)
}

func (heap *Heap) MaxHeapInsert(data int) {
	heap.Data = append(heap.Data, data)
	heap.Size += 1

	index := heap.Size - 1
	for index > 0 && heap.Data[heap.Parent(index)] < heap.Data[index] {
		heap.Data[index], heap.Data[heap.Parent(index)] = heap.Data[heap.Parent(index)], heap.Data[index]
		index = heap.Parent(index)
	}

}

func (heap *Heap) HeapExtractMax() int {
	max := heap.Data[0]

	// Intercambiar primer elemento con el ultimo
	heap.Data[0] = heap.Data[heap.Size-1]

	// Eliminar ultimo elemento del array
	heap.Data = heap.Data[:heap.Size-1]
	heap.Size -= 1

	heap.Data = heap.Max_heapify(0)

	return max
}

func (heap *Heap) HeapMaximum() int {
	return heap.Data[0]
}

func (heap *Heap) Print() {

	var cadena string = "[ "

	// Imprimir el arreglo
	for i := 0; i < heap.Size; i++ {
		cadena += strconv.Itoa(heap.Data[i]) + ", "
	}

	cadena += " ]"

	println(cadena)
}
