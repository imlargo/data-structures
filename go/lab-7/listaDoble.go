package main

type NodoDoble[T any] struct {
	data T
	prev *NodoDoble[T]
	next *NodoDoble[T]
}

func (nodoDoble *NodoDoble[T]) getData() T {
	return nodoDoble.data
}

func (nodoDoble *NodoDoble[T]) setData(data T) {
	nodoDoble.data = data
}

func (nodoDoble *NodoDoble[T]) getNext() *NodoDoble[T] {
	return nodoDoble.next
}

func (nodoDoble *NodoDoble[T]) getPrev() *NodoDoble[T] {
	return nodoDoble.prev
}

func (nodoDoble *NodoDoble[T]) setPrev(prev *NodoDoble[T]) {
	nodoDoble.prev = prev
}

func (nodoDoble *NodoDoble[T]) setNext(next *NodoDoble[T]) {
	nodoDoble.next = next
}

type ListaDoble[T any] struct {
	head *NodoDoble[T]
	tail *NodoDoble[T]
	size int
}

func (listaDoble *ListaDoble[T]) getSize() int {
	return listaDoble.size
}
func (listaDoble *ListaDoble[T]) isEmpty() bool {
	return listaDoble.size == 0
}
func (listaDoble *ListaDoble[T]) first() *NodoDoble[T] {
	return listaDoble.head
}
func (listaDoble *ListaDoble[T]) last() *NodoDoble[T] {
	return listaDoble.tail
}
func (listaDoble *ListaDoble[T]) addFirst(nodo *NodoDoble[T]) {
	if listaDoble.isEmpty() {
		listaDoble.head = nodo
		listaDoble.tail = nodo
	} else {
		nodo.setNext(listaDoble.head)
		listaDoble.head.setPrev(nodo)
		listaDoble.head = nodo
	}
	listaDoble.size++
}
func (listaDoble *ListaDoble[T]) addLast(nodo *NodoDoble[T]) {
	if listaDoble.isEmpty() {
		listaDoble.head = nodo
		listaDoble.tail = nodo
	} else {
		nodo.setPrev(listaDoble.tail)
		listaDoble.tail.setNext(nodo)
		listaDoble.tail = nodo
	}
	listaDoble.size++
}
func (listaDoble *ListaDoble[T]) removeFirst() {
	if listaDoble.isEmpty() {
		return
	}
	listaDoble.head = listaDoble.head.getNext()
	listaDoble.head.setPrev(nil)
	listaDoble.size--
}
func (listaDoble *ListaDoble[T]) removeLast() {
	if listaDoble.isEmpty() {
		return
	}
	listaDoble.tail = listaDoble.tail.getPrev()
	listaDoble.tail.setNext(nil)
	listaDoble.size--
}
