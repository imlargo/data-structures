package main

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
	"data-structures/structures/pila"
)

func main() {
	println("\n---> Problema #1 <---\n")
	problema1()
	println("\n---> Problema #2 <---\n")
	problema2()
	println("\n---> Problema #3 <---\n")
	problema3()
}

func problema1() {

	println("Pila con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método pop() de forma sucesiva")

	var pila pila.Pila[int] = pila.Pila[int]{Lista: &lista_simple.ListaSimple[int]{}}

	for n := 2; n <= 10; n += 2 {
		pila.Push(n)
	}

	for !pila.IsEmpty() {
		println(pila.Pop())
	}
}

func problema2() {
	println("Cola con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método dequeue() de forma sucesiva.")

	var cola cola.Cola[int] = cola.Cola[int]{Lista: &lista_simple.ListaSimple[int]{}}

	for n := 2; n <= 10; n += 2 {
		cola.Enqueue(n)
	}

	for !cola.IsEmpty() {
		println(cola.Dequeue())
	}
}

func problema3() {

	// Ingrese 5 usuarios en la cola e invoque el método toFile() o Invoque el método atenderSiguiente() dos veces y vuelva a llamar el método toFile()
	var turnoUsuario TurnoUsuario = TurnoUsuario{
		registro:          &cola.Cola[User]{Lista: &lista_simple.ListaSimple[User]{}},
		usuariosAtendidos: &pila.Pila[User]{Lista: &lista_simple.ListaSimple[User]{}},
	}

	turnoUsuario.registrar(User{nombre: "Juan", id: "001"})
	turnoUsuario.registrar(User{nombre: "Aleja", id: "002"})
	turnoUsuario.registrar(User{nombre: "Jose", id: "003"})
	turnoUsuario.registrar(User{nombre: "Sebas", id: "004"})
	turnoUsuario.registrar(User{nombre: "Aura", id: "005"})
	turnoUsuario.registrar(User{nombre: "Sara", id: "005"})

	turnoUsuario.toFile()

	// Recorrer lista simple de usuarios pendientes
	println("Usuarios pendientes: ", turnoUsuario.registro.GetSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.GetSize())

	println("\nAtendiendo 2 usuarios...\n")
	turnoUsuario.atenderSiguiente()
	turnoUsuario.atenderSiguiente()

	println("Usuarios pendientes: ", turnoUsuario.registro.GetSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.GetSize())

	turnoUsuario.toFile()
}
