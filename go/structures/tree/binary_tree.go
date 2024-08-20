package tree

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
)

type Node[T any] struct {
	Left  *Node[T]
	Right *Node[T]
	Data  T
}

func (nodo *Node[T]) hasLeft() bool {
	return nodo.Left != nil
}
func (nodo *Node[T]) hasRight() bool {
	return nodo.Right != nil
}

func (nodo *Node[T]) isInternal() bool {
	return nodo.hasLeft() || nodo.hasRight()
}

type BinaryTree[T any] struct {
	Root *Node[T]
	Size int
}

func (tree *BinaryTree[T]) IsEmpty() bool {
	return tree.Size == 0
}

func (tree *BinaryTree[T]) IsRoot(nodo *Node[T]) bool {
	return tree.Root == nodo
}

func (tree *BinaryTree[T]) Parent(nodo *Node[T]) *Node[T] {

	if tree.IsRoot(nodo) {
		return nil
	}

	cola := cola.Cola[*Node[T]]{Lista: &lista_simple.ListaSimple[*Node[T]]{}}
	cola.Enqueue(tree.Root)

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
