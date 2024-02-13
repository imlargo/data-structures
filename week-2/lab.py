class Usuario:
    def __init__(self, nombre = None, idUsuario = None):
        self.nombre = nombre
        self.id = idUsuario
        
        self.fecha_nacimiento = None
        self.ciudad_nacimiento = None
        self.tel = None
        self.email = None
        self.dir = None
    
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
        return f'{self.nombre}, {self.id}, {self.fecha_nacimiento}, {self.ciudad_nacimiento}, {self.tel}, {self.email}, {self.dir}'
    
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
    def __init__(self, calle = None, nomenclatura = None, barrio = None, ciudad = None, edificio = None, apto = None):
        self.calle = calle
        self.nomenclatura = nomenclatura
        self.barrio = barrio
        self.ciudad = ciudad
        self.edificio = edificio
        self.apto = apto
    
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

def app_1():
    fecha = Fecha(16, 7, 2005)
    direccion = Direccion()

    direccion.setCalle('Calle 1')
    direccion.setNomenclatura('Cra 1')
    direccion.setBarrio('Barrio 1')
    direccion.setCiudad('Ciudad 1')
    direccion.setEdificio('Edificio 1')
    direccion.setApto('Apto 1')

    print(
        direccion
    )

    usuario = Usuario('Juan', 123)
    usuario.setFecha_nacimiento(fecha)
    usuario.setDir(direccion)
    usuario.setEmail('imlargo')
    usuario.setTel('1234567890')
    usuario.setCiudad_nacimiento('Valledupar')

    print(
        usuario
    )
    
def app_2():
    
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
        input('Ingrese la calle de su direccion: ')
    )
    direccion.setBarrio(
        input('Ingrese la calle de su direccion: ')
    )
    direccion.setCiudad(
        input('Ingrese la calle de su direccion: ')
    )
    direccion.setEdificio(
        input('Ingrese la calle de su direccion: ')
    )
    direccion.setApto(
        input('Ingrese la calle de su direccion: ')
    )

    print(
        direccion
    )

    print("Bienvenido! ingrese los datos para registrar un usuario")
    
    nombre = input('Ingrese su nombre: ')
    idUsuario = input('Ingrese su id: ')
    usuario = Usuario(
        nombre,
        idUsuario
    )
    usuario.setFecha_nacimiento(fecha)
    usuario.setDir(direccion)
    usuario.setEmail(
        input('Ingrese su email: ')
    )
    usuario.setTel(
        input('Ingrese su telefono: ')
    )
    usuario.setCiudad_nacimiento(
        input('Ingrese su ciudad de nacimiento: ')
    )
    pass