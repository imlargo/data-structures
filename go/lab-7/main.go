package main

func main() {
	problema1()
	problema2()
	problema3()
}

func problema1() {

	println("Pila con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método pop() de forma sucesiva")

	var pila Pila[int] = Pila[int]{lista: &ListaSimple[int]{}}
	pila.push(2)
	pila.push(4)
	pila.push(6)
	pila.push(8)
	pila.push(10)

	for !pila.isEmpty() {
		println(pila.pop())
	}
}

func problema2() {
	println("Cola con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método dequeue() de forma sucesiva.")

	var cola Cola[int] = Cola[int]{lista: &ListaSimple[int]{}}

	cola.enqueue(2)
	cola.enqueue(4)
	cola.enqueue(6)
	cola.enqueue(8)
	cola.enqueue(10)

	for !cola.isEmpty() {
		println(cola.dequeue())
	}
}

func problema3() {

	println("Turno de usuarios: ")

	// Ingrese 5 usuarios en la cola e invoque el método toFile() o Invoque el método atenderSiguiente() dos veces y vuelva a llamar el método toFile()
	var turnoUsuario TurnoUsuario = TurnoUsuario{
		registro:          &Cola[User]{lista: &ListaSimple[User]{}},
		usuariosAtendidos: &Pila[User]{lista: &ListaSimple[User]{}},
	}

	turnoUsuario.registrar(User{nombre: "Juan", id: "001"})
	turnoUsuario.registrar(User{nombre: "Aleja", id: "002"})
	turnoUsuario.registrar(User{nombre: "Jose", id: "003"})
	turnoUsuario.registrar(User{nombre: "Valdi", id: "004"})
	turnoUsuario.registrar(User{nombre: "Chicha", id: "005"})
	turnoUsuario.registrar(User{nombre: "Alejo", id: "005"})

	turnoUsuario.toFile()
	// Recorrer lista simple de usuarios pendientes
	println("Usuarios pendientes: ", turnoUsuario.registro.getSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.getSize())

	turnoUsuario.atenderSiguiente()
	turnoUsuario.atenderSiguiente()

	println("Usuarios pendientes: ", turnoUsuario.registro.getSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.getSize())

	turnoUsuario.toFile()
}
