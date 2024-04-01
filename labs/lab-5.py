from typing import List
import time

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

        usuario = Usuario(
            input('Ingrese su nombre: '),
            input('Ingrese su id: ')
        )
        usuario.setCiudad_nacimiento(
            input('Ingrese su ciudad de nacimiento: ')
        )
        usuario.setTel(
            input('Ingrese su numero detelefono: ')
        )
        usuario.setEmail(input('Ingrese su correo electronico: '))  

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

        usuario.setFecha_nacimiento(fecha)
        usuario.setDir(direccion)
        return usuario

class Agenda:

    def __init__(self, capacidad: int = None):
        self._registro: List[Usuario] = [ None for x in range(capacidad) ]
        self._no_reg: int = 0
        self._capacidad: int = capacidad
        pass

    def agregar(self, usuario: Usuario) -> bool:
        inAgenda = self.buscar(usuario.getId()) != -1
        if (not inAgenda) and (self._no_reg < self._capacidad):
            self._registro[self._no_reg] = usuario
            self._no_reg += 1
            return True
        else:
            return False
        pass

    def buscar(self, userId: str) -> int:
        for i in range(self._no_reg):
            user = self._registro[i]
            if user == None:
                return -1
            if user.getId() == userId:
                return i
        return -1

    def eliminar(self, userId: str) -> bool:
        index = self.buscar(userId)
        if index == -1:
            return False

        for i in range(index + 1, self._no_reg):
            self._registro[i-1] = self._registro[i]

        self._registro[self._no_reg-1] = None
        self._no_reg -= 1
        return True

    def toFile(self) -> None:
        texto = ""
        for i in range(self._no_reg):
            texto += f"{self._registro[i].__str__()}\n" if i < self._no_reg - 1 else f"{self._registro[i].__str__()}"
        
        file = open('agenda.txt', 'w')
        file.write(texto)
        file.close()
        pass

    def importar(self):
        file = open('agenda.txt', 'r')        
        datos = file.read().split('\n')
        file.close()
        
        for informacion in datos:
            usuario = Utils.convertStringToUser(informacion)
            self.agregar(usuario)
        pass

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
    pass

class Mensaje:
    def __init__(self, titulo, contenido, remitente, destinatario):
        self._titulo = titulo
        self._contenido = contenido
        self._remitente = remitente
        self._destinatario = destinatario
        
        # Automaticamente asignar fecha
        self._fecha = time.strftime('%d/%m/%Y - %H:%M:%S')
    pass

class Bandeja:
    
    def __init__(self, usuario):
        self._usuario = usuario
        self._mensajes = DoubleList()
        pass
    
    def addMensaje(self, mensaje):
        self._mensajes.addLast(mensaje)
        pass
    
    def mostrarMensajes(self):
        i = 1
        node = self._mensajes.first()
        while node != None:
            mensaje = node.getData()
            print(f"{i}. Fecha: {mensaje._fecha}, Titulo: {mensaje._titulo}, Remitente: {mensaje._remitente.getNombre()}")
            i += 1
            node = node.getNext()
        pass
    
    def showMensaje(self, n):
        node = self._mensajes.first()
        
        if n == 1:
            mensaje = node.getData()
            print(f"Fecha: {mensaje._fecha}\nTitulo: {mensaje._titulo}\nContenido: {mensaje._contenido}")
            return
            
        for i in range(n-1):
            node = node.getNext()

        mensaje = node.getData()
        print(f"Fecha: {mensaje._fecha}\nTitulo: {mensaje._titulo}\nContenido: {mensaje._contenido}")
        pass
    
    def enviarMensaje(self, titulo, mensaje, destinatario):

        mensaje = Mensaje(titulo, mensaje, self._usuario, destinatario)
        destinatario._bandeja.addMensaje(mensaje)
        pass
    
    pass

class GestionUsuarios:
    
    def __init__(self):
        self._empleados = self.cargarEmpleados()
        pass
    
    def cargarEmpleados(self):
        fileEmpleados = open('Empleados.txt', 'r')
        filePasswords = open('Password.txt', 'r')
        
        datos = zip(fileEmpleados.read().split('\n'), filePasswords.read().split('\n'))
        
        empleados = DoubleList()
        
        for i in len(datos):
            infoUser, infoPassword = datos[i]
            
            empleado = Utils.convertStringToUser(infoUser)
            datosPassword = infoPassword.split(',')
            
            # userId = datosPassword[0]
            empleado.setPassword(datosPassword[1])
            empleado.setRol(datosPassword[2])
            
            if i == 0:
                empleados.addFirst(empleado)
                continue
            
            empleados.addLast(empleado)
            
        fileEmpleados.close()
        filePasswords.close()
            
        return empleados
    
    def buscarEmpleado(self, userId):
        node = self._empleados.first()
        while node != None:
            empleado = node.getData()
            if empleado.getId() == userId:
                return empleado
            node = node.getNext()
        
        return None

class Sistema:
    
    def __init__(self):
        self.gestionUsuarios = GestionUsuarios()
        self.currentUser = self.iniciarSesion()
        
        print(f'Bienvenido al sistema {self.currentUser.getNombre()}')
        print(f"Rol: {self.currentUser.getRol()}")
        
        # Mostrar opciones
        if self.currentUser.isAdmin:
            self.menuAdmin()
        else:
            self.menuEmpleado()
        
        pass
    
    def menuAdmin(self):
        print('Menu de administrador')
        print('1. Registrar empleado')
        
        pass
    
    def menuEmpleado(self):
        print('Menu de empleado')
        print('1. Enviar mensaje')
        pass
    
    def menuMensaje(self):
        documento = input('Ingrese el documento del destinatario: ')
        
        titulo = input('Ingrese el titulo del mensaje: ')
        contenido = input('Ingrese el contenido del mensaje: ')
        
        destinatario = self.gestionUsuarios.buscarEmpleado(documento)
        
        # Enviar el mensaje
        self.currentUser._bandeja.enviarMensaje(titulo, contenido, destinatario)
        pass
    
    def leerMensajes(self):
        print('Bandeja de entrada:')
        self.currentUser._bandeja.mostrarMensajes()
        
        option = input('Desea leer un mensaje? (s/n): ')
        if option == 's':
            n = int(input('Ingrese el numero del mensaje: '))
            self.currentUser._bandeja.showMensaje(n)
            pass
        else:
            return None
        
        pass
    
    def iniciarSesion(self):
        documento = int(input('Ingrese su documento: '))
        password = input('Ingrese su contraseña: ')
        
        usuario = self.gestionUsuarios.buscarEmpleado(documento)
        
        if usuario == None:
            print('El usuario no existe')
            return None
        
        return usuario if usuario.getPassword() == password else None
    
    def registrarEmpleado(self):
        pass
    
    pass

def main():
    sistema = Sistema()
    pass