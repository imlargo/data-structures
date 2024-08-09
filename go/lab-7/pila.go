package main

type Pila[T any] struct {
	lista *ListaSimple[T]
}

func (pila *Pila[T]) getSize() int {
	return pila.lista.getSize()
}
func (pila *Pila[T]) isEmpty() bool {
	return pila.lista.isEmpty()
}
func (pila *Pila[T]) push(data T) {
	nodo := NodoSimple[T]{data: data}
	pila.lista.addFirst(&nodo)
}
func (pila *Pila[T]) pop() T {
	data := pila.lista.removeFirst()
	return data
}
func (pila *Pila[T]) top() T {
	return pila.lista.first().getData()
}
