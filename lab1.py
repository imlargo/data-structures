#Programa 1 
#   Variables, expresiones y estructuras de control: Diseñe un programa que le pida al usuario
#   ingresar n números enteros, donde n también es ingresado por el usuario. El programa debe calcular (i) el valor
#   máximo, (ii) el valor mínimo, (iii) la suma de los enteros ingresados y (iv) el valor promedio.

def calcularNumeros():
    n = int(input("Ingrese la cantidad de números a ingresar: "))
    datos = [
        int(
            input(f"Ingrese el número {i+1}: ")
        ) for i in range(n)
    ]
    maximo = datos[0]
    minimo = datos[0]
    total = 0
    for num in datos:
        total += num
        if num > maximo:
            maximo = num
        if num < minimo:
            minimo = num
    promedio = total / n
    
    return {
        "maximo": maximo,
        "minimo": minimo,
        "total": total,
        "promedio": promedio
    }

#Programa 2
#   Archivos y cadenas: Diseñe un programa que lea el archivo de texto test_pr2.txt (archivo adjunto a
#   esta guía) y cuente cuantas veces se repite la palabra “en” en el párrafo leído. Tenga en cuenta que el párrafo esta
#   dividido en varias líneas. (No modifique el archivo de texto).

def contarPalabra():
    texto = open("test_pr2.txt", "r").read()
    palabras = texto.split(" ")

    total = 0

    for palabra in palabras:
        if palabra == "en":
            total += 1
    return total


#Programa 3
#   Tablas y diccionarios: Diseñe un programa que realice el control de acceso a una aplicación. El
#   programa almacena los usuarios y contraseñas mostrados en la siguiente tabla. El programa solicita al usuario
#   ingresar su nombre y contraseña. Si el usuario introduce de forma correcta los datos, el programa muestra en
#   pantalla el mensaje “Acceso permitido”, en caso contrario indicara “Datos incorrectos”. Si el usuario se equivoca tres
#   veces consecutivas el programa indicara “Lo siento, su acceso no es permitido”. 

def iniciarSesion():
    
    credenciales = {
        "Juan1223": "J12an*.",
        "Maria2345": "M23a*.",
        "Pablo1459": "P14o*.",
        "Ana3456": "A34a*",
    }
    
    for i in range(3):
        user = input("Ingrese su nombre de usuario: ")
        password = input("Ingrese su contraseña: ")

        isCorrecto = credenciales[user] == password
        mensaje = "Acceso permitido" if isCorrecto else "Datos incorrectos"
        print(mensaje)
        
        if isCorrecto:
            return
        
    print("Lo siento, su acceso no es permitido")