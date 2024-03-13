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
        
    def setNext(self, next):
        self._next = next

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
        
    def setNext(self, next):
        self._next = next
        
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
        self._head = node
        self._size += 1
        pass
    
    def addLast(self, data):
        node = SimpleNode(data)
        self._tail.setNext(node)
        self._tail = node
        self._size += 1
        pass
    
    def removeFirst(self):
        data = self._head.getData()
        self._head = self._head.getNext()
        self._size -= 1
        return data
    

class DoubleList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0
        
    def Size(self):
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
        newNode.setNext(node.getNext())
        newNode.setPrev(node)
        node.setNext(newNode)
        newNode.getNext().setPrev(newNode)
        self._size += 1
        pass

    def addBefore(self, node, data):
        newNode = DoubleNode(data)
        newNode.setPrev(node.getPrev())
        newNode.setNext(node)
        
        node.setPrev(newNode)
        node.getPrev().setNext(newNode)
        self._size += 1
        pass
    
def main_1():
    listaDoble = DoubleList()
    listaDoble.addFirst(2)
    for n in range(4, 21, 2):
        listaDoble.addLast(n)
    
    siguiente = listaDoble.first()
    while siguiente != None:
        print(siguiente.getData())
        siguiente = siguiente.getNext()
        
    # Eliminar los números 1, 10 y 20
    listaDoble.removeFirst()
    listaDoble.removeLast()

    print("Iniciando")

    siguiente = listaDoble.first()
    while siguiente != None:
        if siguiente.getData() == 10:
            listaDoble.remove(siguiente)
            break
        siguiente = siguiente.getNext()


    siguiente = listaDoble.first()
    while siguiente != None:
        print(siguiente.getData())
        siguiente = siguiente.getNext()

def main_2():
    pass

main_1()