from typing import List
import time
import os
import sys

class Usuario:
    def __init__(self, nombre: str = None, idUsuario: str = None):
        self.nombre = nombre
        self.id = idUsuario
        self.fecha_nacimiento: Fecha = None
        self.ciudad_nacimiento: str = None
        self.tel: str = None
        self.email: str = None
        self.dir: Direccion = None
    
    # Setter methods
    def setNombre(self, nombre):
        self.nombre = nombre
    def setId(self, id):
        self.id = int(id)
    def setFecha_nacimiento(self, fecha):
        self.fecha_nacimiento = fecha
    def setCiudad_nacimiento(self, ciudad):
        self.ciudad_nacimiento = ciudad
    def setTel(self, telefono):
        self.tel = telefono
    def setEmail(self, email):
        self.email = email
    def setDir(self, dir):
        self.dir = dir
    
    # Getter methods
    def getNombre(self):
        return self.nombre
    
    def getId(self):
        return self.id
    
    def getFecha_nacimiento(self):
        return self.fecha_nacimiento
    
    def getCiudad_nacimiento(self):
        return self.ciudad_nacimiento
    
    def getTel(self):
        return self.tel
    
    def getEmail(self):
        return self.email
    
    def getDir(self):
        return self.dir
    
    # String representation
    def __str__(self):
        return f'{self.nombre} - {self.id} - {self.fecha_nacimiento} - {self.ciudad_nacimiento} - {self.tel} - {self.email} - {self.dir}'
    
class Fecha:
    def __init__(self, dd = None, mm = None, aa = None):
        self.dd = dd
        self.mm = mm
        self.aa = aa
        pass

    def setDia(self, dd):
        self.dd = dd
    def setMes(self, mm):
        self.mm = mm
    def setA(self, aa):
        self.aa = aa

    def getDia(self):
        return self.dd
    def getMes(self):
        return self.mm
    def getA(self):
        return self.aa
    
    def __str__(self):
        return f'{self.dd}/{self.mm}/{self.aa}'
    
class Direccion:
    def __init__(self, calle: str = None, nomenclatura: str = None, barrio: str = None, ciudad: str = None, edificio: str = None, apto: str = None):
        self.calle: str = calle
        self.nomenclatura: str = nomenclatura
        self.barrio: str = barrio
        self.ciudad: str = ciudad
        self.edificio: str = edificio
        self.apto: str = apto
    
    def setCalle(self, calle):
        self.calle = calle
    def setNomenclatura(self, n):
        self.nomenclatura = n
    def setBarrio(self, b):
        self.barrio = b
    def setCiudad(self, ci):
        self.ciudad = ci
    def setEdificio(self, e):
        self.edificio = e
    def setApto(self, a):
        self.apto = a
    def getCalle(self):
        return self.calle
    def getNomenclatura(self):
        return self.nomenclatura
    def getBarrio(self):
        return self.barrio
    def getCiudad(self):
        return self.ciudad
    def getEdificio(self):
        return self.edificio
    def getApto(self):
        return self.apto
    
    def __str__(self):
        return f'{self.calle}, {self.nomenclatura}, {self.barrio}, {self.ciudad}, {self.edificio}, {self.apto}'

class Utils:
    
    predefinidos = [
        "Juan - 123 - 16/7/2005 - Valledupar - 123456789 - imlargo@unal - Calle 6, Cra 6, Barrio Robledo, Ciudad 1, Edificio 1, Apto 100",
        "Alejandra - 201 - 16/7/2005 - Valledupar - 987654321 - aleja@unal - Calle 5, Cra 5, Barrio 1, Ciudad 1, Edificio 2, Apto 201",
        "Harrison - 302 - 16/7/2005 - Valledupar - 12349876 - harri@unal - Calle 4, Cra 4, Barrio 1, Ciudad 1, Edificio 3, Apto 302",
        "JulianMoreno - 403 - 16/7/2005 - Valledupar - 98712340 - juliangod@unal - Calle 3, Cra 3, Barrio 1, Ciudad 1, Edificio 4, Apto 403",
        "TaylorSwift - 504 - 16/7/2005 - Valledupar - 19283745 - taylor@unal - Calle 2, Cra 2, Barrio 1, Ciudad 1, Edificio 5, Apto 504",
    ]

    @classmethod
    def convertStringToUser(cls, informacion: str) -> Usuario:
        
        # Estrucutra de la informacion de usuario: nombre, id, fecha_nacimiento, ciudad_nacimiento, tel, email, dir
        datosUsuario = informacion.split(' - ')
        
        # Estructura de la fecha dd/mm/aa
        fecha = cls.convertStringToDate(datosUsuario[2])
        
        # Estructura de la direccion: calle, nomenclatura, barrio, ciudad, edificio, apto
        direccion = cls.convertStringToDireccion(datosUsuario[6])
        
        empleado = Empleado(datosUsuario[0], int(datosUsuario[1]))
        empleado.setFecha_nacimiento(fecha)
        empleado.setCiudad_nacimiento(datosUsuario[3])
        empleado.setTel(datosUsuario[4])
        empleado.setEmail(datosUsuario[5])
        empleado.setDir(direccion)
        return empleado
    
    @classmethod
    def convertStringToDate(cls, fecha: str) -> Fecha:
        # Separar la fecha en dia, mes y año
        datosFecha = fecha.split('/')
        return Fecha(int(datosFecha[0]), int(datosFecha[1]), int(datosFecha[2]))
    
    @classmethod
    def convertStringToDireccion(cls, direccion: str) -> Direccion:
        # Separar la direccion en sus componentes: calle, nomenclatura, barrio, ciudad, edificio, apto
        datosDireccion = direccion.split(', ')
        return Direccion(datosDireccion[0], datosDireccion[1], datosDireccion[2], datosDireccion[3], datosDireccion[4], datosDireccion[5])

    @classmethod
    def createFromInput(cls):

        empleado = Empleado(
            input('Ingrese su nombre: '),
            int(input('Ingrese su id: '))
        )
        empleado.setCiudad_nacimiento(
            input('Ingrese su ciudad de nacimiento: ')
        )
        empleado.setTel(
            input('Ingrese su numero detelefono: ')
        )
        empleado.setEmail(input('Ingrese su correo electronico: '))  

        # Fecha
        fecha = Fecha(
            int(input('Ingrese el dia de su fecha de nacimiento: ')),
            int(input('Ingrese el mes de su fecha de nacimiento: ')),
            int(input('Ingrese el año de su fecha de nacimiento: '))
        )

        # Direccion
        direccion = Direccion()

        print("Informacion de direccion:")
        direccion.setCalle(
            input('Ingrese la calle de su direccion: ')
        )
        direccion.setNomenclatura(
            input('Ingrese la nomenclatura de su direccion: ')
        )
        direccion.setBarrio(
            input('Ingrese el barrio de su direccion: ')
        )
        direccion.setCiudad(
            input('Ingrese la ciudad de su direccion: ')
        )
        direccion.setEdificio(
            input('Ingrese el edificio de su direccion: ')
        )
        direccion.setApto(
            input('Ingrese el apartamento de su direccion: ')
        )

        empleado.setFecha_nacimiento(fecha)
        empleado.setDir(direccion)
        return empleado

#....

class SimpleNode:
    def __init__(self, data = None):
        self.data = data
        self._next = None
        
    # Getters and setters
    def getData(self):
        return self.data
    
    def getNext(self):
        return self._next
    
    def setData(self, data):
        self.data = data
        
    def setNext(self, nextNode):
        self._next = nextNode

class DoubleNode:
    def __init__(self, data = None):
        self.data = data
        self._next = None
        self._prev = None
        
    # Getters and setters
    def getData(self):
        return self.data
    
    def setData(self, data):
        self.data = data
    
    def getNext(self):
        return self._next
    
    def getPrev(self):
        return self._prev
        
    def setNext(self, nextNode):
        self._next = nextNode
        
    def setPrev(self, prev):
        self._prev = prev

class SimpleList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0
        
    def size(self):
        return self._size
    
    def isEmpty(self):
        return self._size == 0
    
    def setSize(self, size):
        self._size = size

    def first(self):
        return self._head
    
    def last(self):
        return self._tail

    def addFirst(self, data):
        node = SimpleNode(data)
        node.setNext(self._head)
        if self.isEmpty():
            self._tail = node
        self._head = node
        self._size += 1
        pass
    
    def addLast(self, data):
        node = SimpleNode(data)
        if self.isEmpty():
            self._tail = node
            self._head = node
        else:
            self._tail.setNext(node)
            self._tail = node
        self._size += 1
        pass
    
    def removeFirst(self):
        data = self._head.getData()
        self._head = self._head.getNext()
        self._size -= 1
        return data

    def removeLast(self):
        if self.size() == 1:
            return self.removeFirst()
        else:
            anterior = self._head
            while anterior.getNext() != self._tail:
                anterior = anterior.getNext()
            
            data = self._tail.getData()
            anterior.setNext(None)
            self._tail = anterior
            self._size -= 1
            return data
    
    def printData(self):
        node = self._head
        while node != None:
            print(node.getData())
            node = node.getNext()
        pass

class DoubleList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0
        
    def size(self):
        return self._size
    
    def isEmpty(self):
        return self._size == 0
    def first(self):
        return self._head
    def last(self):
        return self._tail
    
    def addFirst(self, data):
        node = DoubleNode(data)
        node.setNext(self._head)        
        if self.isEmpty():
            self._tail = node
        else:
            self._head.setPrev(node)
        self._head = node
        self._size += 1
        pass
    
    def addLast(self, data):
        
        if self.isEmpty():
            return self.addFirst(data)
        
        node = DoubleNode(data)
        node.setPrev(self._tail)
        self._tail.setNext(node)
        self._tail = node
        self._size += 1
        pass
        
    def removeFirst(self):
        data = self._head.getData()
        self._head = self._head.getNext()
        self._head.setPrev(None)
        self._size -= 1
        return data
    
    def removeLast(self):
        data = self._tail.getData()
        self._tail = self._tail.getPrev()
        self._tail.setNext(None)
        self._size -= 1
        return data

    def remove(self, node):
        
        if node == self._head:
            return self.removeFirst()
        
        if node == self._tail:
            return self.removeLast()
        
        data = node.getData()

        prev = node.getPrev()
        nextNode = node.getNext()
        prev.setNext(nextNode)
        nextNode.setPrev(prev)

        self._size -= 1
        return data
        
    def addAfter(self, node, data):
        newNode = DoubleNode(data)
        newNode.setPrev(node)
        newNode.setNext(node.getNext())

        node.getNext().setPrev(newNode)
        node.setNext(newNode)
        self._size += 1
        pass

    def addBefore(self, node, data):
        newNode = DoubleNode(data)

        newNode.setNext(node)
        newNode.setPrev(node.getPrev())

        node.getPrev().setNext(newNode)
        node.setPrev(newNode)
        self._size += 1
        pass

    def printData(self):
        node = self._head
        while node != None:
            print(node.getData())
            node = node.getNext()
        pass


class Empleado(Usuario):

    def __init__(self, nombre: str = None, idUsuario: str = None):
        super().__init__(nombre, idUsuario)
        
        self._password: str = None
        self._rol: str = None
        
        self.isAdmin = None
        self.gestionUsuarios = None
        self._bandeja = None
        pass
    
    def initBandeja(self):
        self._bandeja = Bandeja(self)
        pass
    
    def getPassword(self):
        return self._password
    def getRol(self):
        return self._rol
    def setPassword(self, password):
        self._password = password
    def setRol(self, rol):
        self._rol = rol
        self.isAdmin = (rol == 'administrador')
    def setGestionUsuarios(self, gestionUsuarios):
        self.gestionUsuarios = gestionUsuarios
    pass

class Mensaje:
    def __init__(self, titulo, contenido, remitente, destinatario, fecha):
        self._titulo = titulo
        self._contenido = contenido
        self._remitente = remitente
        self._destinatario = destinatario
        self._fecha = fecha
        
    def __str__(self) -> str:
        return f"{self._titulo} ; {self._contenido} ; {self._remitente.getId()} ; {self._destinatario.getId()} ; {self._fecha}"
    
    @classmethod
    def fromString(cls, informacion: str, remitente: Empleado):
        datosMensaje = informacion.split(' ; ')
        destinatario = remitente.gestionUsuarios.buscarEmpleado(datosMensaje[3])
        return cls(datosMensaje[0], datosMensaje[1], remitente, destinatario, datosMensaje[4])
    pass

class Bandeja:
    
    def __init__(self, usuario):
        self._usuario = usuario
        self._mensajes = DoubleList()

        self.importar()
        pass
    
    def addMensaje(self, mensaje):
        self._mensajes.addLast(mensaje)
        self.export()
        pass
    
    def mostrarMensajes(self):
        i = 1
        node = self._mensajes.first()
        while node != None:
            mensaje = node.getData()
            print(f"    {i}. Fecha: {mensaje._fecha}, Titulo: {mensaje._titulo}, Remitente: {mensaje._remitente.getNombre()}")
            i += 1
            node = node.getNext()
        pass
    
    def showMensaje(self, n):
        node = self._mensajes.first()
        
        if n == 1:
            mensaje = node.getData()
            print(f"    Fecha: {mensaje._fecha}\n    Titulo: {mensaje._titulo}\n    Contenido: {mensaje._contenido}")
            return
            
        for i in range(n-1):
            node = node.getNext()

        mensaje = node.getData()
        print(f"    Fecha: {mensaje._fecha}\n    Titulo: {mensaje._titulo}\n    Contenido: {mensaje._contenido}")
        pass
    
    
    def borrarMensaje(self, n):
        if n == 1:
            self._mensajes.removeFirst()
            self.export()
            return
        
        node = self._mensajes.first()
        for i in range(n-1):
            node = node.getNext()
        
        self._mensajes.remove(node)
        self.export()
        pass
    
    
    def enviarMensaje(self, titulo, mensaje, destinatario):

        # Automaticamente asignar fecha
        mensaje = Mensaje(
            titulo,
            mensaje,
            self._usuario,
            destinatario,
            time.strftime('%d/%m/%Y - %H:%M:%S')
        )
        destinatario._bandeja.addMensaje(mensaje)
        pass
    
    def export(self):
        
        data = []
        
        node = self._mensajes.first()
        while node != None:
            mensaje = node.getData()
            data.append(mensaje.__str__())
            node = node.getNext()

        file = open(f'{self._usuario.getId()}BA.txt', 'w')
        file.write('\n'.join(data))
        file.close()
        pass
    
    def importar(self):
        
        # Comprobar si el archivo existe
        if not os.path.exists(f'{self._usuario.getId()}BA.txt'):
            return None
        
        file = open(f'{self._usuario.getId()}BA.txt', 'r')
        datos = file.read().split('\n')
        file.close()
        
        for informacion in datos:
            mensaje = Mensaje.fromString(informacion, self._usuario)
            self._mensajes.addLast(mensaje)
    pass

class GestionUsuarios:
    
    def __init__(self):
        self._empleados = self.cargarEmpleados()
        self.initBandejas()
        pass
    
    def addEmpleado(self, empleado):
        self._empleados.addLast(empleado)
        self.sortEmpleados()
        self.guardarEmpleados()
        pass
    
    def eliminarEmpleado(self, userId):
        nodo = self.buscarNodoEmpleado(int(userId))
        
        if nodo == None:
            return False

        self._empleados.remove(nodo)
        self.guardarEmpleados()

        return True
    
    def sortEmpleados(self):
        empleados = []
        
        # Convertir la lista en un arreglo
        node = self._empleados.first()
        while node != None:
            empleados.append(node.getData())
            node = node.getNext()
            
        # Ordenamiento de burbuja
        for i in range(len(empleados)):
            for j in range(len(empleados)-1):
                if empleados[j].getId() > empleados[j+1].getId():
                    empleados[j], empleados[j+1] = empleados[j+1], empleados[j]
                    
                    
        # Convertir arreglo en lista enlazada
        listaEmpleados = DoubleList()
        [listaEmpleados.addLast(empleado) for empleado in empleados]
            
        self._empleados = listaEmpleados
        pass
    
    def guardarEmpleados(self):
        
        fileEmpleados = open('Empleados.txt', 'w')
        filePasswords = open('Password.txt', 'w')
        
        dataEmpleados = []
        dataPasswords = []
        
        node = self._empleados.first()
        while node != None:
            empleado = node.getData()
            dataEmpleados.append(f'{empleado.__str__()}')
            dataPasswords.append(f'{empleado.getId()},{empleado.getPassword()},{empleado.getRol()}')
            node = node.getNext()
            
        fileEmpleados.write('\n'.join(dataEmpleados))
        filePasswords.write('\n'.join(dataPasswords))
        
        fileEmpleados.close()
        filePasswords.close()
        pass
    
    def cargarEmpleados(self):
        fileEmpleados = open('Empleados.txt', 'r')
        filePasswords = open('Password.txt', 'r')
        
        dataEmpleados = fileEmpleados.read().split('\n')
        dataPasswords = filePasswords.read().split('\n')
        
        empleados = DoubleList()
        
        for infoUser, infoPassword in zip(dataEmpleados, dataPasswords):
            
            empleado = Utils.convertStringToUser(infoUser)
            datosPassword = infoPassword.split(',')
            
            # userId = datosPassword[0]
            empleado.setPassword(datosPassword[1])
            empleado.setRol(datosPassword[2])
            
            empleado.setGestionUsuarios(self)
            empleados.addLast(empleado)
            
        fileEmpleados.close()
        filePasswords.close()

        return empleados
    
    def initBandejas(self):
        node = self._empleados.first()
        while node != None:
            empleado = node.getData()
            empleado.initBandeja()
            node = node.getNext()
        pass
    
    def buscarEmpleado(self, userId):
        node = self._empleados.first()
        
        while node != None:
            empleado = node.getData()
            if empleado.getId() == int(userId):
                return empleado
            node = node.getNext()
        
        return None
    
    def buscarNodoEmpleado(self, userId):
        node = self._empleados.first()
        
        while node != None:
            if node.getData().getId() == userId:
                return node
            node = node.getNext()
        return None

class Sistema:
    
    def __init__(self):
        self.gestionUsuarios = GestionUsuarios()
        
        while True:
            self.currentUser = self.iniciarSesion()
            
            if self.currentUser == None:
                print('Usuario o contraseña incorrectos')
                continue
            
            print(f'Bienvenido al sistema {self.currentUser.getNombre()}')
            print(f"Rol: {self.currentUser.getRol()}")

            # Mostrar opciones
            if self.currentUser.isAdmin:
                self.menuAdmin()
            else:
                self.menuEmpleado()
        pass
    
    def menuAdmin(self):
        while True:
            print('--- Menu de administrador ---')
            print('1. Enviar mensaje')
            print('2. Revisar bandeja de entrada')
            print('3. Registrar empleado')
            print('4. Eliminar empleado')
            print('5. Cambiar contraseña')

            print('6. Cerrar sesion')
            print('7. Salir')

            print(". . . . . . . . . . ")

            option = input('Ingrese la opcion: ')
            match option:
                case '1':
                    self.menuMensaje()
                case '2':
                    self.revisarBandeja()
                case '3':
                    self.registrarEmpleado()
                case '4':
                    self.eliminarEmpleado()
                case '5':
                    self.cambiarPassword()
                case '6':
                    return
                case '7':
                    print('Saliendo del sistema, hasta luego!')
                    sys.exit()
                case _:
                    print('Opcion no valida')
        pass
    
    def menuEmpleado(self):
        while True:
            print('--- Menu de empleado ---')
            print('1. Enviar mensaje')
            print('2. Revisar bandeja de entrada')
            print('3. Cerrar sesion')
            print('4. Salir')

            print(". . . . . . . . . . ")

            option = input('Ingrese la opcion: ')
            print()

            match option:
                case '1':
                    self.menuMensaje()
                case '2':
                    self.revisarBandeja()
                case '3':
                    return
                case '4':
                    print('Saliendo del sistema, hasta luego!')
                    exit()
                case _:
                    print('Opcion no valida')
            pass
    
    def menuMensaje(self):
        print('--- Enviar mensaje ---')
        
        documento = input('Ingrese el documento del destinatario: ')
        titulo = input('Ingrese el titulo del mensaje: ')
        contenido = input('Ingrese el contenido del mensaje: ')
        
        destinatario = self.gestionUsuarios.buscarEmpleado(documento)

        if not destinatario:
            print(f'\n    > No existe ningun empleado con documento {documento}\n')
            return
        
        # Enviar el mensaje
        self.currentUser._bandeja.enviarMensaje(
            titulo,
            contenido,
            destinatario,
        )
        
        print(f'\n    > Mensaje enviado a {destinatario.getNombre()}\n')
        pass
    
    def revisarBandeja(self):
        print('--- Bandeja de entrada ---\n')
        self.currentUser._bandeja.mostrarMensajes()
        print(". . . . . . . . . .")
        
        print("1. Leer mensaje")
        print("2. Borrar mensaje")
        print("3. Salir")
        
        option = input('Ingrese la opcion: ')
        
        match option:
            case '1':
                n = int(input('Ingrese el numero del mensaje: '))
                print()
                self.currentUser._bandeja.showMensaje(n)
                print()
            case '2':
                n = int(input('Ingrese el numero del mensaje: '))
                self.currentUser._bandeja.borrarMensaje(n)
                print("\nMensaje borrado con exito\n")
            case '3':
                return
            case _:
                print('Opcion no valida')
        pass
    
    def iniciarSesion(self):
        documento = int(input('Ingrese su documento: '))
        password = input('Ingrese su contraseña: ')
        
        usuario = self.gestionUsuarios.buscarEmpleado(documento)
        
        if usuario == None:
            print('El usuario no existe')
            return None
        
        if usuario.getPassword() == password:
            return usuario
        
        return None
    
    def registrarEmpleado(self):
        empleado = Utils.createFromInput()
        print()
        
        password = input('Ingrese la contraseña: ')
        empleado.setPassword(password)
        
        rol = input('Ingrese el rol: ')
        empleado.setRol(rol)
        
        empleado.setGestionUsuarios(self.gestionUsuarios)
        self.gestionUsuarios.addEmpleado(empleado)
        empleado.initBandeja()

        print("Empleado registrado con exito!")
        pass
    
    def eliminarEmpleado(self):
        documento = input('Ingrese el documento del usuario: ')
        response = self.gestionUsuarios.eliminarEmpleado(documento)
        
        print(
            "Usuario eliminado con exito" if response == True else "El usuario no existe"
        )
        pass
    
    def cambiarPassword(self):
        documento = input('Ingrese el documento del usuario: ')
        
        empleado = self.gestionUsuarios.buscarEmpleado(documento)
        if empleado == None:
            print('El usuario no existe')
            return
        
        password = input('Ingrese la nueva contraseña: ')
        empleado.setPassword(password)
        self.gestionUsuarios.sortEmpleados()
        self.gestionUsuarios.guardarEmpleados()

        pass
    
    pass

Sistema()