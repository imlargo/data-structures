package main

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