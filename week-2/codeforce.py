# 1. Acumulados al inverso
n = int(input())
numeros = input().split(" ")

acum = int(numeros[-1])
for i in range(n - 2, -1, -1):
    acum += int(numeros[i])
    print(acum)


# 2. Juego trisolariano
n = int(input())
turnos = [
    input().split(", ")
    for i in range(3)
]

pSo = 0
pLar = 0
pIs = 0

for i in range(n):
    turnoSo = int(turnos[0][i])
    turnoLar = int(turnos[1][i])
    turnoIs = int(turnos[2][i])

    if (turnoSo + turnoLar + turnoIs) % 2 == 0:
        if turnoSo % 2 == 0:
            pSo += 1
        if turnoLar % 2 == 0:
            pLar += 1
        if turnoIs % 2 == 0:
            pIs += 1
    else:
        if turnoSo % 2 != 0:
            pSo += 1 
        if turnoLar % 2 != 0:
            pLar += 1 
        if turnoIs % 2 != 0:
            pIs += 1 

print(f"SO:{pSo}, LAR:{pLar}, IS:{pIs}")

# 3. Criptografía para dummies
for x in range(int(input())):
    mensaje = input().replace(" ", "")
    string = ""
    lastIndex = len(mensaje) - 1
    if ((lastIndex+1) % 2 == 0):
        for i in range(0, lastIndex, 2):
            string += mensaje[lastIndex-(i+1)] + mensaje[lastIndex-i]
    else:
        for i in range(0, lastIndex - 1, 2):
            string += mensaje[lastIndex-(i+1)] + mensaje[lastIndex-i]
        string += mensaje[0]
    print(string)


# 4. Recorrido arcoíris
n = int(input())
chars = input().split(", ")
string = ""
lastIndex = n-1
for i in range(n//2):
    string += chars[i] + chars[lastIndex-i]
string += chars[n//2] if n % 2 != 0 else ""
print(string)

# 5. Arte callejero
n = int(input())
for x in range(n):
    datos = input().split(" ") # Integrantes y monedas
    denominaciones =  input().split(" ")

    maximo = int(denominaciones[0])
    minimo = maximo
    
    valores = {}

    for i in range(int(datos[1])):
        valor = int(denominaciones[i])
        modulo = str(i % int(datos[0]))
        valores[modulo] = valores[modulo] + valor if (modulo in valores) else valor
        acum = valores[modulo]
        if acum > maximo:
            maximo = acum
        elif acum < minimo:
            minimo = acum
        
# 4. Arte callejero
n = int(input())
for x in range(n):
    datos = input().split(" ")
    denominaciones =  input().split(" ")
    valores = {}
    for i in range(int(datos[1])):
        modulo = str(i % int(datos[0]))
        valores[modulo] = valores[modulo] + int(denominaciones[i]) if (modulo in valores) else int(denominaciones[i])
    maximo = valores["0"]
    minimo = maximo
    for k in valores.values():
        if k > maximo:
            maximo = k
        elif k < minimo:
            minimo = k
    print(maximo - minimo)


# 6. Serpientes y escaleras
n = int(input())
for x in range(n):
    size = int(input())
    tablero = input().split(" ")
    recorridos = []

    i = int(tablero[0])
    recorridos.append(i)

    while (-1 < i < len(tablero) or (i not in recorridos)):
        i = tableros[
            tableros[i]
        ]
        pass