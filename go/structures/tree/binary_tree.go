package tree

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
)

type BinaryTree[T any] struct {
	Root *Node[T]
	Size int
}

func (tree *BinaryTree[T]) getRoot() *Node[T] {
	return tree.Root
}

func (tree *BinaryTree[T]) IsEmpty() bool {
	return tree.Size == 0
}

func (tree *BinaryTree[T]) IsRoot(nodo *Node[T]) bool {
	return tree.getRoot() == nodo
}

func (tree *BinaryTree[T]) Parent(nodo *Node[T]) *Node[T] {

	if tree.IsRoot(nodo) {
		return nil
	}

	cola := cola.Cola[*Node[T]]{Lista: &lista_simple.ListaSimple[*Node[T]]{}}
	cola.Enqueue(tree.getRoot())

	var temp *Node[T]

	for !cola.IsEmpty() {
		temp = cola.Dequeue()
		if temp.Left == nodo || temp.Right == nodo {
			return temp
		}
		if temp.hasLeft() {
			cola.Enqueue(temp.Left)
		}
		if temp.hasRight() {
			cola.Enqueue(temp.Right)
		}
	}

	return nil
}
func (tree *BinaryTree[T]) Depth(nodo *Node[T]) int {

	if tree.IsRoot(nodo) {
		return 0
	}

	return 1 + tree.Depth(tree.Parent(nodo))
}

func (tree *BinaryTree[T]) Height(nodo *Node[T]) int {
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

func (tree *BinaryTree[T]) AddRoot(nodo *Node[T]) {
	tree.Root = nodo
	tree.Size = 1
}
func (tree *BinaryTree[T]) InsertLeft(nodo1 *Node[T], nodo2 *Node[T]) {
	nodo1.Left = nodo2
	tree.Size++
}

func (tree *BinaryTree[T]) InsertRight(nodo1 *Node[T], nodo2 *Node[T]) {
	nodo1.Right = nodo2
	tree.Size++
}

func (tree *BinaryTree[T]) Remove(nodo *Node[T]) {

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
		nodo = nil

		tree.Size--
		return
	}

	// Si es hoja
	if parent.Left == nodo {
		parent.Left = nil
	} else {
		parent.Right = nil
	}

	nodo = nil
	tree.Size--

}

func (tree *BinaryTree[T]) Min(nodo *Node[T]) T {
	if nodo.hasLeft() {
		return tree.Min(nodo.Left)
	} else {
		return nodo.Data
	}
}

func (tree *BinaryTree[T]) Max(nodo *Node[T]) T {
	if nodo.hasRight() {
		return tree.Max(nodo.Right)
	} else {
		return nodo.Data
	}
}
