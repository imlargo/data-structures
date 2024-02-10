# 1. Encuentra el cuadrado

n = int(input())
print(n*n)

# 2. Hola mundo bilingüe

for x in range(int(input())):
    print("Hola mundo" if (x % 2 == 0) else "Hello world")

# 3. Positivos y negativos
sumPos = 0
sumNeg = 0
for i in range(int(input())):
    num = int(input())
    if (num > 0):
        sumPos += num
    else:
        sumNeg += num
print(f"positivos {sumPos}, negativos {sumNeg}")

# 4. Los 4 fantásticos
n = int(input())
for d in [2, 3, 5, 7]:
    if n % d == 0:
        print(f"es multiplo de {d}")
        exit()
print("no es multiplo de ninguno de los primeros cuatro primos")


# 5. Potencias menores
n = int(input())
limit = int(input())
for p in range(1, limit):
    r = n ** p
    if r > limit:
        break
    print(r)

# 6. Conjetura de Collatz
n = int(input())
while n > 1:
    print(
        int(n)
    )
    n = n / 2 if (n % 2 == 0) else (3 * n) + 1
print(int(n))
