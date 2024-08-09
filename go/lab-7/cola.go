package main

type Cola[T any] struct {
	lista *ListaSimple[T]
}

func (cola *Cola[T]) getSize() int {
	return cola.lista.getSize()
}
func (cola *Cola[T]) isEmpty() bool {
	return cola.lista.isEmpty()
}
func (cola *Cola[T]) enqueue(data T) {
	var nodo NodoSimple[T] = NodoSimple[T]{data: data}
	cola.lista.addLast(&nodo)
}
func (cola *Cola[T]) dequeue() T {
	data := cola.lista.removeFirst()
	return data
}
func (cola *Cola[T]) first() T {
	return cola.lista.first().getData()
}
