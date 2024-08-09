from typing import List

class Usuario:
    def __init__(self, nombre: str = None, idUsuario: str = None):
        self.nombre = nombre
        self.id = idUsuario
        
    # Setter methods
    def setNombre(self, nombre):
        self.nombre = nombre
    def setId(self, id):
        self.id = int(id)
    
    # Getter methods
    def getNombre(self):
        return self.nombre
    
    def getId(self):
        return self.id
    
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
        
    def setNext(self, nextNode):
        self._next = nextNode

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
        if self.isEmpty():
            self._tail = node
        self._head = node
        self._size += 1
        pass
    
    def addLast(self, data):
        node = SimpleNode(data)
        if self.isEmpty():
            self._tail = node
            self._head = node
        else:
            self._tail.setNext(node)
            self._tail = node
        self._size += 1
        pass
    
    def removeFirst(self):
        data = self._head.getData()
        self._head = self._head.getNext()
        self._size -= 1
        return data

    def removeLast(self):
        if self.size() == 1:
            return self.removeFirst()
        else:
            anterior = self._head
            while anterior.getNext() != self._tail:
                anterior = anterior.getNext()
            
            data = self._tail.getData()
            anterior.setNext(None)
            self._tail = anterior
            self._size -= 1
            return data
    
    def printData(self):
        node = self._head
        while node != None:
            print(node.getData())
            node = node.getNext()
        pass


class Stack:
    def __init__(self):
        self._data = SimpleList()
        
        
    def size(self):
        return self._data.size()


def main():
    
    pass