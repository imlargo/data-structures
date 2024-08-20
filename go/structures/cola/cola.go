package cola

import "data-structures/structures/lista_simple"

type Cola[T any] struct {
	Lista *lista_simple.ListaSimple[T]
}

func (cola *Cola[T]) GetSize() int {
	return cola.Lista.GetSize()
}
func (cola *Cola[T]) IsEmpty() bool {
	return cola.Lista.IsEmpty()
}
func (cola *Cola[T]) Enqueue(data T) {
	var nodo lista_simple.NodoSimple[T] = lista_simple.NodoSimple[T]{Data: data}
	cola.Lista.AddLast(&nodo)
}
func (cola *Cola[T]) Dequeue() T {
	data := cola.Lista.RemoveFirst()
	return data
}
func (cola *Cola[T]) First() T {
	return cola.Lista.First().GetData()
}
