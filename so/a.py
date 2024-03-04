from threading import Thread
from threading import get_ident
from time import sleep

suma = 0
def sumatoria(inferior, superior):
    global suma
    temp = 0
    for i in range(inferior, superior):
        temp += i
    suma += temp
    
def s(a,b):
    i = 0
    for x in range(a,b+1):
        i+= x
    return i
        
    
a = 50
b = 100
mid = a + (b - a) // 2

t2 = Thread(target=sumatoria, args=(a, mid))
t1 = Thread(target=sumatoria, args=(mid, b+1))

t1.start()
t2.start()

t1.join()
t2.join()

print(suma)
print(s(a,b))