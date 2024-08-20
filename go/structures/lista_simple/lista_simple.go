package lista_simple

type NodoSimple[T any] struct {
	Data T
	Next *NodoSimple[T]
}

func (nodoSimple *NodoSimple[T]) GetData() T {
	return nodoSimple.Data
}

func (nodoSimple *NodoSimple[T]) GetNext() *NodoSimple[T] {
	return nodoSimple.Next
}

func (nodoSimple *NodoSimple[T]) SetData(data T) {
	nodoSimple.Data = data
}

func (nodoSimple *NodoSimple[T]) SetNext(next *NodoSimple[T]) {
	nodoSimple.Next = next
}

type ListaSimple[T any] struct {
	Head *NodoSimple[T]
	Tail *NodoSimple[T]
	Size int
}

func (listaSimple *ListaSimple[T]) GetSize() int {
	return listaSimple.Size
}

func (listaSimple *ListaSimple[T]) IsEmpty() bool {
	return listaSimple.Size == 0
}

func (listaSimple *ListaSimple[T]) First() *NodoSimple[T] {
	return listaSimple.Head
}

func (listaSimple *ListaSimple[T]) Last() *NodoSimple[T] {
	return listaSimple.Tail
}

func (listaSimple *ListaSimple[T]) AddFirst(nodo *NodoSimple[T]) {
	if listaSimple.IsEmpty() {
		listaSimple.Head = nodo
		listaSimple.Tail = nodo
	} else {
		nodo.SetNext(listaSimple.Head)
		listaSimple.Head = nodo
	}
	listaSimple.Size++
}

func (listaSimple *ListaSimple[T]) AddLast(nodo *NodoSimple[T]) {
	if listaSimple.IsEmpty() {
		listaSimple.Head = nodo
		listaSimple.Tail = nodo
	} else {
		listaSimple.Tail.SetNext(nodo)
		listaSimple.Tail = nodo
	}
	listaSimple.Size++
}

func (listaSimple *ListaSimple[T]) RemoveFirst() T {
	if !listaSimple.IsEmpty() {
		data := listaSimple.Head.GetData()
		listaSimple.Head = listaSimple.Head.GetNext()
		listaSimple.Size--
		return data
	}
	var zeroValue T
	return zeroValue
}

func (listaSimple *ListaSimple[T]) RemoveLast() T {
	if !listaSimple.IsEmpty() {
		if listaSimple.Size == 1 {

			data := listaSimple.Head.GetData()

			listaSimple.Head = nil
			listaSimple.Tail = nil
			listaSimple.Size--

			return data
		} else {
			data := listaSimple.Tail.GetData()

			current := listaSimple.Head
			for current.GetNext() != listaSimple.Tail {
				current = current.GetNext()
			}
			current.SetNext(nil)
			listaSimple.Tail = current
			listaSimple.Size--

			return data
		}
	}

	var zeroValue T
	return zeroValue
}
