package main

import (
	"data-structures/structures/cola"
	"data-structures/structures/pila"
	"fmt"
	"os"
)

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
