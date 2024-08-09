package main

type NodoSimple[T any] struct {
	data T
	next *NodoSimple[T]
}

func (nodoSimple *NodoSimple[T]) getData() T {
	return nodoSimple.data
}

func (nodoSimple *NodoSimple[T]) getNext() *NodoSimple[T] {
	return nodoSimple.next
}

func (nodoSimple *NodoSimple[T]) setData(data T) {
	nodoSimple.data = data
}

func (nodoSimple *NodoSimple[T]) setNext(next *NodoSimple[T]) {
	nodoSimple.next = next
}

type ListaSimple[T any] struct {
	head *NodoSimple[T]
	tail *NodoSimple[T]
	size int
}

func (listaSimple *ListaSimple[T]) getSize() int {
	return listaSimple.size
}

func (listaSimple *ListaSimple[T]) isEmpty() bool {
	return listaSimple.size == 0
}

func (listaSimple *ListaSimple[T]) first() *NodoSimple[T] {
	return listaSimple.head
}

func (listaSimple *ListaSimple[T]) last() *NodoSimple[T] {
	return listaSimple.tail
}

func (listaSimple *ListaSimple[T]) addFirst(nodo *NodoSimple[T]) {
	if listaSimple.isEmpty() {
		listaSimple.head = nodo
		listaSimple.tail = nodo
	} else {
		nodo.setNext(listaSimple.head)
		listaSimple.head = nodo
	}
	listaSimple.size++
}

func (listaSimple *ListaSimple[T]) addLast(nodo *NodoSimple[T]) {
	if listaSimple.isEmpty() {
		listaSimple.head = nodo
		listaSimple.tail = nodo
	} else {
		listaSimple.tail.setNext(nodo)
		listaSimple.tail = nodo
	}
	listaSimple.size++
}

func (listaSimple *ListaSimple[T]) removeFirst() T {
	if !listaSimple.isEmpty() {
		data := listaSimple.head.getData()
		listaSimple.head = listaSimple.head.getNext()
		listaSimple.size--
		return data
	}
	var zeroValue T
	return zeroValue
}

func (listaSimple *ListaSimple[T]) removeLast() T {
	if !listaSimple.isEmpty() {
		if listaSimple.size == 1 {

			data := listaSimple.head.getData()

			listaSimple.head = nil
			listaSimple.tail = nil
			listaSimple.size--

			return data
		} else {
			data := listaSimple.tail.getData()

			current := listaSimple.head
			for current.getNext() != listaSimple.tail {
				current = current.getNext()
			}
			current.setNext(nil)
			listaSimple.tail = current
			listaSimple.size--

			return data
		}
	}

	var zeroValue T
	return zeroValue
}
