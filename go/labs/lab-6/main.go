package main

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
	"data-structures/structures/pila"
	"fmt"
	"os"
)

func main() {
	problema1()
	problema2()
	problema3()
}

func problema1() {

	println("Pila con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método pop() de forma sucesiva")

	var pila pila.Pila[int] = pila.Pila[int]{Lista: &lista_simple.ListaSimple[int]{}}
	pila.Push(2)
	pila.Push(4)
	pila.Push(6)
	pila.Push(8)
	pila.Push(10)

	for !pila.IsEmpty() {
		println(pila.Pop())
	}
}

func problema2() {
	println("Cola con los valores enteros 2-4-6-8-10, imprima en pantalla la pila haciendo uso del método dequeue() de forma sucesiva.")

	var cola cola.Cola[int] = cola.Cola[int]{Lista: &lista_simple.ListaSimple[int]{}}

	cola.Enqueue(2)
	cola.Enqueue(4)
	cola.Enqueue(6)
	cola.Enqueue(8)
	cola.Enqueue(10)

	for !cola.IsEmpty() {
		println(cola.Dequeue())
	}
}

func problema3() {

	println("Turno de usuarios: ")

	// Ingrese 5 usuarios en la cola e invoque el método toFile() o Invoque el método atenderSiguiente() dos veces y vuelva a llamar el método toFile()
	var turnoUsuario TurnoUsuario = TurnoUsuario{
		registro:          &cola.Cola[User]{Lista: &lista_simple.ListaSimple[User]{}},
		usuariosAtendidos: &pila.Pila[User]{Lista: &lista_simple.ListaSimple[User]{}},
	}

	turnoUsuario.registrar(User{nombre: "Juan", id: "001"})
	turnoUsuario.registrar(User{nombre: "Aleja", id: "002"})
	turnoUsuario.registrar(User{nombre: "Jose", id: "003"})
	turnoUsuario.registrar(User{nombre: "Valdi", id: "004"})
	turnoUsuario.registrar(User{nombre: "Chicha", id: "005"})
	turnoUsuario.registrar(User{nombre: "Alejo", id: "005"})

	turnoUsuario.toFile()
	// Recorrer lista simple de usuarios pendientes
	println("Usuarios pendientes: ", turnoUsuario.registro.GetSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.GetSize())

	turnoUsuario.atenderSiguiente()
	turnoUsuario.atenderSiguiente()

	println("Usuarios pendientes: ", turnoUsuario.registro.GetSize())
	println("Usuarios atendidos: ", turnoUsuario.usuariosAtendidos.GetSize())

	turnoUsuario.toFile()
}

type TurnoUsuario struct {
	registro          *cola.Cola[User]
	usuariosAtendidos *pila.Pila[User]
}

func (turno *TurnoUsuario) registrar(usuario User) {
	turno.registro.Enqueue(usuario)
}

func (turno *TurnoUsuario) atenderSiguiente() {
	siguiente := turno.registro.Dequeue()
	turno.usuariosAtendidos.Push(siguiente)
}

func crearArchivo(nombre string, contenido string) {
	f, err := os.Create(nombre)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(contenido)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (turno *TurnoUsuario) toFile() {
	// almacena los usuarios que quedaron en la cola en un archivo de texto usuariospendientes.txt y los que fueron atendidos en un archivo de texto usuariosatendidos.txt.

	var usuariosPendientes string
	// Recorrer lista simple de usuarios pendientes
	for nodo := turno.registro.Lista.First(); nodo != nil; nodo = nodo.GetNext() {
		var usuario User = nodo.GetData()
		usuariosPendientes += usuario.serializar() + "\n"
	}

	var usuariosAtendidos string
	for nodo := turno.usuariosAtendidos.Lista.First(); nodo != nil; nodo = nodo.GetNext() {
		var usuario User = nodo.GetData()
		usuariosAtendidos += usuario.serializar() + "\n"
	}

	crearArchivo("usuariospendientes.txt", usuariosPendientes)
	crearArchivo("usuariosatendidos.txt", usuariosAtendidos)
}

type User struct {
	nombre string
	id     string
}

func (user *User) serializar() string {
	return user.nombre + " " + user.id
}

func (user *User) deserializar(serializacion string) {
	user.nombre = serializacion[:len(serializacion)-3]
	user.id = serializacion[len(serializacion)-3:]
}
