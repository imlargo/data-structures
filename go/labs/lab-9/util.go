package main

func getUsuarios() []*Usuario {
	var usuarios []*Usuario = make([]*Usuario, 6)
	usuarios[0] = &Usuario{Nombre: "Juan", Id: 1}
	usuarios[1] = &Usuario{Nombre: "Pablo", Id: 2}
	usuarios[2] = &Usuario{Nombre: "Maria", Id: 3}
	usuarios[3] = &Usuario{Nombre: "Ana", Id: 4}
	usuarios[4] = &Usuario{Nombre: "Diana", Id: 5}
	usuarios[5] = &Usuario{Nombre: "Mateo", Id: 6}

	return usuarios
}
