package lista_doble

type NodoDoble[T any] struct {
	Data T
	Prev *NodoDoble[T]
	Next *NodoDoble[T]
}

func (nodoDoble *NodoDoble[T]) GetData() T {
	return nodoDoble.Data
}

func (nodoDoble *NodoDoble[T]) SetData(data T) {
	nodoDoble.Data = data
}

func (nodoDoble *NodoDoble[T]) GetNext() *NodoDoble[T] {
	return nodoDoble.Next
}

func (nodoDoble *NodoDoble[T]) GetPrev() *NodoDoble[T] {
	return nodoDoble.Prev
}

func (nodoDoble *NodoDoble[T]) SetPrev(prev *NodoDoble[T]) {
	nodoDoble.Prev = prev
}

func (nodoDoble *NodoDoble[T]) SetNext(next *NodoDoble[T]) {
	nodoDoble.Next = next
}

type ListaDoble[T any] struct {
	Head *NodoDoble[T]
	Tail *NodoDoble[T]
	Size int
}

func (listaDoble *ListaDoble[T]) GetSize() int {
	return listaDoble.Size
}
func (listaDoble *ListaDoble[T]) IsEmpty() bool {
	return listaDoble.Size == 0
}
func (listaDoble *ListaDoble[T]) First() *NodoDoble[T] {
	return listaDoble.Head
}
func (listaDoble *ListaDoble[T]) Last() *NodoDoble[T] {
	return listaDoble.Tail
}
func (listaDoble *ListaDoble[T]) AddFirst(nodo *NodoDoble[T]) {
	if listaDoble.IsEmpty() {
		listaDoble.Head = nodo
		listaDoble.Tail = nodo
	} else {
		nodo.SetNext(listaDoble.Head)
		listaDoble.Head.SetPrev(nodo)
		listaDoble.Head = nodo
	}
	listaDoble.Size++
}
func (listaDoble *ListaDoble[T]) AddLast(nodo *NodoDoble[T]) {
	if listaDoble.IsEmpty() {
		listaDoble.Head = nodo
		listaDoble.Tail = nodo
	} else {
		nodo.SetPrev(listaDoble.Tail)
		listaDoble.Tail.SetNext(nodo)
		listaDoble.Tail = nodo
	}
	listaDoble.Size++
}
func (listaDoble *ListaDoble[T]) RemoveFirst() {
	if listaDoble.IsEmpty() {
		return
	}
	listaDoble.Head = listaDoble.Head.GetNext()
	if listaDoble.Head != nil {
		listaDoble.Head.SetPrev(nil)
	} else {
		listaDoble.Tail = nil
	}

	listaDoble.Size--
}
func (listaDoble *ListaDoble[T]) RemoveLast() {
	if listaDoble.IsEmpty() {
		return
	}
	listaDoble.Tail = listaDoble.Tail.GetPrev()
	if listaDoble.Tail != nil {
		listaDoble.Tail.SetNext(nil)
	} else {
		listaDoble.Head = nil
	}

	listaDoble.Size--
}
func (listaDoble *ListaDoble[T]) Remove(nodo *NodoDoble[T]) T {
	data := nodo.GetData()

	if nodo.Prev == nil {
		listaDoble.RemoveFirst()
		return data
	}

	if nodo.Next == nil {
		listaDoble.RemoveLast()
		return data
	}

	prev := nodo.GetPrev()
	nextNode := nodo.GetNext()
	prev.SetNext(nextNode)
	nextNode.SetPrev(prev)
	listaDoble.Size--
	return data
}
