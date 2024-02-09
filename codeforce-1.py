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