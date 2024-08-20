package pila

import "data-structures/structures/lista_simple"

type Pila[T any] struct {
	Lista *lista_simple.ListaSimple[T]
}

func (pila *Pila[T]) GetSize() int {
	return pila.Lista.GetSize()
}
func (pila *Pila[T]) IsEmpty() bool {
	return pila.Lista.IsEmpty()
}
func (pila *Pila[T]) Push(data T) {
	nodo := lista_simple.NodoSimple[T]{Data: data}
	pila.Lista.AddFirst(&nodo)
}
func (pila *Pila[T]) Pop() T {
	data := pila.Lista.RemoveFirst()
	return data
}
func (pila *Pila[T]) Top() T {
	return pila.Lista.First().GetData()
}
