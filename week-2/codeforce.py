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