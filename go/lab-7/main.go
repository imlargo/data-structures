package main

func main() {
	// Usar lista simple
	listaSimple := ListaSimple[int]{}
	listaSimple.addFirst(&NodoSimple[int]{data: 1})
	listaSimple.addLast(&NodoSimple[int]{data: 2})
	listaSimple.addLast(&NodoSimple[int]{data: 3})
	listaSimple.addLast(&NodoSimple[int]{data: 4})
	listaSimple.addLast(&NodoSimple[int]{data: 5})

	// Recorrer lista simple
	for nodo := listaSimple.first(); nodo != nil; nodo = nodo.getNext() {
		println(nodo.getData())
	}

	listaSimple.removeFirst()
	listaSimple.removeLast()

	println("Eliminados")

	// Recorrer lista simple
	for nodo := listaSimple.first(); nodo != nil; nodo = nodo.getNext() {
		println(nodo.getData())
	}

}
