from typing import List

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
        self.id = id
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
        return f'{self.nombre} - {self.id} - {self.fecha_nacimiento} - {self.ciudad_nacimiento} - {self.tel} - {self.email}, {self.dir}'
    
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
        "Juan, 123, 16/7/2005, Valledupar, 123456789, imlargo@unal, Calle 6, Cra 6, Barrio Robledo, Ciudad 1, Edificio 1, Apto 100",
        "Alejandra, 201, 16/7/2005, Valledupar, 987654321, aleja@unal, Calle 5, Cra 5, Barrio 1, Ciudad 1, Edificio 2, Apto 201",
        "Harrison, 302, 16/7/2005, Valledupar, 12349876, harri@unal, Calle 4, Cra 4, Barrio 1, Ciudad 1, Edificio 3, Apto 302",
        "JulianMoreno, 403, 16/7/2005, Valledupar, 98712340, juliangod@unal, Calle 3, Cra 3, Barrio 1, Ciudad 1, Edificio 4, Apto 403",
        "TaylorSwift, 504, 16/7/2005, Valledupar, 19283745, imlargo@unal, Calle 2, Cra 2, Barrio 1, Ciudad 1, Edificio 5, Apto 504",
        "Juan, 605, 16/7/2005, Valledupar, 13131313, imlargo@unal, Calle 1, Cra 1, Barrio 1, Ciudad 1, Edificio 6, Apto 605",
    ]

    @classmethod
    def convertStringToUser(cls, informacion: str) -> Usuario:
        
        # Estrucutra de la informacion de usuario: nombre, id, fecha_nacimiento, ciudad_nacimiento, tel, email, dir
        datosUsuario = informacion.split(' - ')
        
        # Estructura de la fecha dd/mm/aa
        fecha = cls.convertStringToDate(datosUsuario[2])
        
        # Estructura de la direccion: calle, nomenclatura, barrio, ciudad, edificio, apto
        direccion = cls.convertStringToDireccion(datosUsuario[6])
        
        usuario = Usuario(datosUsuario[0], datosUsuario[1])
        usuario.setFecha_nacimiento(fecha)
        usuario.setCiudad_nacimiento(datosUsuario[3])
        usuario.setTel(datosUsuario[4])
        usuario.setEmail(datosUsuario[5])
        usuario.setDir(direccion)
        return usuario
    
    @classmethod
    def convertStringToDate(cls, fecha: str) -> Fecha:
        # Separar la fecha en dia, mes y año
        datosFecha = fecha.split('/')
        return Fecha(datosFecha[0], datosFecha[1], datosFecha[2])
    
    @classmethod
    def convertStringToDireccion(cls, direccion: str) -> Direccion:
        # Separar la direccion en sus componentes: calle, nomenclatura, barrio, ciudad, edificio, apto
        datosDireccion = direccion.split(', ')
        return Direccion(datosDireccion[0], datosDireccion[1], datosDireccion[2], datosDireccion[3], datosDireccion[4], datosDireccion[5])

class Agenda:

    def __init__(self, capacidad: int = None):
        self._registro: List[Usuario] = []
        self._no_reg: int = 0
        self._capacidad: int = capacidad
        pass
    
    def agregar(self, usuario: Usuario) -> bool:
        inAgenda = self.buscar(usuario.getId()) != -1
        if (not inAgenda):
            self._registro.append(usuario)
            self._no_reg += 1
            return True
        else:
            return False
        pass
    
    def buscar(self, userId: str) -> int:
        for i in range(len(self._registro)):
            usuario = self._registro[i]
            if  usuario.getId() == userId:
                return i
        return -1
        
    def eliminar(self, userId: str) -> bool:
        index = self.buscar(userId)
        
        if index == -1:
            return False

        self._registro.pop(index)
        self._no_reg -= 1
        return True
    
    def toFile(self) -> None:
        texto = "\n".join(
            map(lambda x: x.__str__(), self._registro)
        )
        file = open('agenda.txt', 'w')
        file.write(texto)
        file.close()
        pass

    def importar(self):
        file = open('agenda.txt', 'r')
        file.close()
        
        datos = file.read().split('\n')
        
        for informacion in datos:
            usuario = Utils.convertStringToUser(informacion)
            self.agregar(usuario)
        pass


def main_1():
    # Inicializar 5 usuarios en una agenda con toda su información
    agenda = Agenda(5)

    for userString in Utils.predefinidos:
        user = Utils.convertStringToUser(userString)
        agenda.agregar(user)

    # Buscar un usuario por su número de id e imprimir la posición
    print(
        agenda.buscar(123)
    )

    # Guardar los usuarios de la agenda en un archivo con el metodo toFile() 
    agenda.toFile()
    pass

def main_2():
    # Importar la agenda
    agenda = Agenda(5)
    agenda.importar()

    # Imprimir en pantalla los 5 usuarios leídos
    for usuario in agenda._registro:
        print(usuario)
        pass

    # Eliminar un usuario dado su número de id
    agenda.eliminar(123)

    # Guardar los usuarios actualizados de la agenda con toFile()
    agenda.toFile()
    pass