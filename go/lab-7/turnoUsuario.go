package main

import (
	"fmt"
	"os"
)

type TurnoUsuario struct {
	registro          *Cola[User]
	usuariosAtendidos *Pila[User]
}

func (turno *TurnoUsuario) registrar(usuario User) {
	turno.registro.enqueue(usuario)
}

func (turno *TurnoUsuario) atenderSiguiente() {
	siguiente := turno.registro.dequeue()
	turno.usuariosAtendidos.push(siguiente)
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
	for nodo := turno.registro.lista.first(); nodo != nil; nodo = nodo.getNext() {
		var usuario User = nodo.getData()
		usuariosPendientes += usuario.serializar() + "\n"
	}

	var usuariosAtendidos string
	for nodo := turno.usuariosAtendidos.lista.first(); nodo != nil; nodo = nodo.getNext() {
		var usuario User = nodo.getData()
		usuariosAtendidos += usuario.serializar() + "\n"
	}

	crearArchivo("usuariospendientes.txt", usuariosPendientes)
	crearArchivo("usuariosatendidos.txt", usuariosAtendidos)
}
