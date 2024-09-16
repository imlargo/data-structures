package hash_table

import (
	"data-structures/structures/lista_doble"
)

type ChainedHashTable[T any] struct {
	Array    []*lista_doble.ListaDoble[*Entry[T]]
	Size     int
	HashFunc func(int, int) int
}

func New[T any](size int, hashFunc func(int, int) int) *ChainedHashTable[T] {

	array := make([]*lista_doble.ListaDoble[*Entry[T]], size)

	for i := 0; i < size; i++ {
		lista := lista_doble.ListaDoble[*Entry[T]]{}
		array[i] = &lista
	}

	chainedHashTable := ChainedHashTable[T]{
		Array:    array,
		Size:     size,
		HashFunc: hashFunc,
	}

	return &chainedHashTable
}

func (hashTable *ChainedHashTable[T]) Insert(key int, value T) {
	index := hashTable.HashFunc(key, hashTable.Size)
	entry := Entry[T]{Key: key, Value: value}
	nodo := lista_doble.NodoDoble[*Entry[T]]{Data: &entry}
	hashTable.Array[index].AddLast(&nodo)
}

func (hashtable *ChainedHashTable[T]) Search(key int) (T, bool) {
	index := hashtable.HashFunc(key, hashtable.Size)

	lista := hashtable.Array[index]
	nodo := lista.First()
	for nodo != nil {
		entry := nodo.GetData()
		if entry.Key == key {
			return entry.Value, true
		}
		nodo = nodo.GetNext()
	}

	var zeroValue T
	return zeroValue, false
}

func (hashtable *ChainedHashTable[T]) Delete(key int) {
	index := hashtable.HashFunc(key, hashtable.Size)

	lista := hashtable.Array[index]
	nodo := lista.First()
	for nodo != nil {
		entry := nodo.GetData()
		if entry.Key == key {
			lista.Remove(nodo)
			break
		}
		nodo = nodo.GetNext()
	}
}
