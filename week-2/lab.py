class Usuario:
    def __init__(self, nombre = None, idUsuario = None):
        self._nombre = nombre
        self._id = idUsuario
        
        self._fecha_nacimiento = None
        self._ciudad_nacimiento = None
        self._tel = None
        self._email = None
        self._dir = None
    
    # Setters
    def setNombre(self, nombre):
        self._nombre = nombre
    def setId(self, id):
        self._id = id
    def setFecha_nacimiento(self, fecha):
        self._fecha_nacimiento = fecha
    def setCiudad_nacimiento(self, ciudad):
        self._ciudad_nacimiento = ciudad
    def setTel(self, telefono):
        self._tel = telefono
    def setEmail(self, email):
        self._email = email
    def setDir(self, dir):
        self._dir = dir
    
    # Getters
    def getNombre(self):
        return self._nombre
    
    def getId(self):
        return self._id
    
    def getFecha_nacimiento(self):
        return self._fecha_nacimiento
    
    def getCiudad_nacimiento(self):
        return self._ciudad_nacimiento
    
    def getTel(self):
        return self._tel
    
    def getEmail(self):
        return self._email
    
    def getDir(self):
        return self._dir
    
    # ToString
    def __str__(self):
        return f'{self._nombre}, {self._id}, {self._fecha_nacimiento}, {self._ciudad_nacimiento}, {self._tel}, {self._email}, {self._dir}'
    

class Fecha:
    def __init__(self, dd = 16, mm = 7, aa = 2005):
        self._dd = dd
        self._mm = mm
        self._aa = aa
        pass

    def setDia(self, dd):
        self._dd = dd
    def setMes(self, mm):
        self._mm = mm
    def setA(self, aa):
        self._aa = aa

    def getDia(self):
        return self._dd
    def getMes(self):
        return self._mm
    def getA(self):
        return self._aa
    
    # ToString
    def __str__(self):
        return f'{self._dd}/{self._mm}/{self._aa}'
    
    
class Direccion:
    def __init__(self):
        self._calle = None
        self._nomenclatura = None
        self._barrio = None
        self._ciudad = None
        self._edificio = None
        self._apto = None
    
    # Setters
    def setCalle(self, calle):
        self._calle = calle
        
    def setNomenclatura(self, n):
        self._nomenclatura = n

    def setBarrio(self, b):
        self._barrio = b

    def setCiudad(self, ci):
        self._ciudad = ci

    def setEdificio(self, e):
        self._edificio = e

    def setApto(self, a):
        self._apto = a

    # Getters
    def getCalle(self):
        return self._calle
    
    def getNomenclatura(self):
        return self._nomenclatura
    
    def getBarrio(self):
        return self._barrio
    
    def getCiudad(self):
        return self._ciudad
    
    def getEdificio(self):
        return self._edificio
    
    def getApto(self):
        return self._apto
    
    # ToString
    def __str__(self):
        return f'{self._calle}, {self._nomenclatura}, {self._barrio}, {self._ciudad}, {self._edificio}, {self._apto}'


def main_1():
    fecha = Fecha(16, 7, 2005)
    print("Fecha:")
    print(
        fecha
    )

    direccion = Direccion()
    direccion.setCalle('Calle 1')
    direccion.setNomenclatura('Cra 1')
    direccion.setBarrio('Barrio 1')
    direccion.setCiudad('Ciudad 1')
    direccion.setEdificio('Edificio 1')
    direccion.setApto('Apto 1')

    print("Direccion:")
    print(
        direccion
    )

    usuario = Usuario('Juan', 123)
    usuario.setFecha_nacimiento(fecha)
    usuario.setDir(direccion)
    usuario.setEmail('imlargo@unal')
    usuario.setTel('1234567890')
    usuario.setCiudad_nacimiento('Valledupar')

    print("Usuario:")
    print(
        usuario
    )
    
def main_2():
    
    print("Bienvenido! ingrese los datos para registrar un usuario")
    
    nombre = input('Ingrese su nombre: ')
    idUsuario = input('Ingrese su id: ')
    usuario = Usuario(
        nombre,
        idUsuario
    )

    usuario.setEmail(
        input('Ingrese su email: ')
    )
    usuario.setTel(
        input('Ingrese su telefono: ')
    )
    usuario.setCiudad_nacimiento(
        input('Ingrese su ciudad de nacimiento: ')
    )

    dia = int(input('Ingrese el dia de su fecha de nacimiento: '))
    mes = int(input('Ingrese el mes de su fecha de nacimiento: '))
    year = int(input('Ingrese el año de su fecha de nacimiento: '))
    fecha = Fecha(
        dia, mes, year
    )

    direccion = Direccion()

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

    print(
        direccion
    )

    pass