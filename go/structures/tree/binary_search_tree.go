package tree

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
)

type BinarySearchTree[T any] struct {
	Root *Node[T]
	Size int
}

func (tree *BinarySearchTree[T]) getRoot() *Node[T] {
	return tree.Root
}

func (tree *BinarySearchTree[T]) IsEmpty() bool {
	return tree.Size == 0
}

func (tree *BinarySearchTree[T]) IsRoot(nodo *Node[T]) bool {
	return tree.getRoot() == nodo
}

func (tree *BinarySearchTree[T]) Parent(nodo *Node[T]) *Node[T] {

	if tree.IsRoot(nodo) {
		return nil
	}

	cola := cola.Cola[*Node[T]]{Lista: &lista_simple.ListaSimple[*Node[T]]{}}
	cola.Enqueue(tree.getRoot())

	temp := tree.Root

	for !cola.IsEmpty() && cola.First().Left != nodo && cola.First().Right != nodo {

		temp := cola.Dequeue()
		if temp.hasLeft() {
			cola.Enqueue(temp.Left)
		}
		if temp.hasRight() {
			cola.Enqueue(temp.Right)
		}

	}

	return temp
}
func (tree *BinarySearchTree[T]) Depth(nodo *Node[T]) int {

	if tree.IsRoot(nodo) {
		return 0
	}

	return 1 + tree.Depth(tree.Parent(nodo))
}

func (tree *BinarySearchTree[T]) Height(nodo *Node[T]) int {
	if !nodo.isInternal() {
		return 0
	}

	var h int = 0

	leftHeight := tree.Height(nodo.Left)
	rightHeight := tree.Height(nodo.Right)

	if leftHeight > rightHeight {
		h = leftHeight
	} else {
		h = rightHeight
	}

	return 1 + h
}

func (tree *BinarySearchTree[T]) AddRoot(nodo *Node[T]) {
	tree.Root = nodo
	tree.Size = 1
}
func (tree *BinarySearchTree[T]) InsertLeft(nodo1 *Node[T], nodo2 *Node[T]) {
	nodo1.Left = nodo2
	tree.Size++
}

func (tree *BinarySearchTree[T]) InsertRight(nodo1 *Node[T], nodo2 *Node[T]) {
	nodo1.Right = nodo2
	tree.Size++
}

func (tree *BinarySearchTree[T]) Remove(nodo *Node[T]) {

	var parent *Node[T] = tree.Parent(nodo)

	// Si tiene al menos un hijo
	if nodo.isInternal() {
		var child *Node[T]

		if nodo.hasLeft() {
			child = nodo.Left
		}
		if nodo.hasRight() {
			child = nodo.Right
		}

		if parent.Left == nodo {
			parent.Left = child
		} else {
			parent.Right = child
		}

		nodo.Left = nil
		nodo.Right = nil

		tree.Size--
		return
	}

	if parent.Left == nodo {
		parent.Left = nil
	} else {
		parent.Right = nil
	}

	nodo = nil
	tree.Size--

}

func (tree *BinarySearchTree[T]) Min(nodo *Node[T]) T {
	if nodo.hasLeft() {
		return tree.Min(nodo.Left)
	} else {
		return nodo.Data
	}
}

func (tree *BinarySearchTree[T]) Max(nodo *Node[T]) T {
	if nodo.hasRight() {
		return tree.Max(nodo.Right)
	} else {
		return nodo.Data
	}
}

// Binary Search Tree
func (tree *BinarySearchTree[T]) Find(key int) {

}

func (tree *BinarySearchTree[T]) Insert(data T, key int) {

}

func (tree *BinarySearchTree[T]) RemoveByKey(key int) *T {
	return nil
}
