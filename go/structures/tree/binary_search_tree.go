package tree

import (
	"data-structures/structures/cola"
	"data-structures/structures/lista_simple"
)

type BinarySearchTree[T any] struct {
	Root *Node[*BST_Entry[T]]
	Size int
}

func (tree *BinarySearchTree[T]) getRoot() *Node[*BST_Entry[T]] {
	return tree.Root
}

func (tree *BinarySearchTree[T]) IsEmpty() bool {
	return tree.Size == 0
}

func (tree *BinarySearchTree[T]) IsRoot(nodo *Node[*BST_Entry[T]]) bool {
	return tree.getRoot() == nodo
}

func (tree *BinarySearchTree[T]) Parent(nodo *Node[*BST_Entry[T]]) *Node[*BST_Entry[T]] {

	if tree.IsRoot(nodo) {
		return nil
	}

	cola := cola.Cola[*Node[*BST_Entry[T]]]{Lista: &lista_simple.ListaSimple[*Node[*BST_Entry[T]]]{}}
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
func (tree *BinarySearchTree[T]) Depth(nodo *Node[*BST_Entry[T]]) int {

	if tree.IsRoot(nodo) {
		return 0
	}

	return 1 + tree.Depth(tree.Parent(nodo))
}

func (tree *BinarySearchTree[T]) Height(nodo *Node[*BST_Entry[T]]) int {
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

func (tree *BinarySearchTree[T]) AddRoot(nodo *Node[*BST_Entry[T]]) {
	tree.Root = nodo
	tree.Size = 1
}
func (tree *BinarySearchTree[T]) InsertLeft(nodo1 *Node[*BST_Entry[T]], nodo2 *Node[*BST_Entry[T]]) {
	nodo1.Left = nodo2
	tree.Size++
}

func (tree *BinarySearchTree[T]) InsertRight(nodo1 *Node[*BST_Entry[T]], nodo2 *Node[*BST_Entry[T]]) {
	nodo1.Right = nodo2
	tree.Size++
}

func (tree *BinarySearchTree[T]) Remove(nodo *Node[*BST_Entry[T]]) {

	var parent *Node[*BST_Entry[T]] = tree.Parent(nodo)

	// Si tiene al menos un hijo
	if nodo.isInternal() {
		var child *Node[*BST_Entry[T]]

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

func (tree *BinarySearchTree[T]) Min(nodo *Node[*BST_Entry[T]]) T {
	if nodo.hasLeft() {
		return tree.Min(nodo.Left)
	} else {
		return nodo.Data.GetData()
	}
}

func (tree *BinarySearchTree[T]) Max(nodo *Node[*BST_Entry[T]]) T {
	if nodo.hasRight() {
		return tree.Max(nodo.Right)
	} else {
		return nodo.Data.GetData()
	}
}

// Binary Search Tree
func (tree *BinarySearchTree[T]) Find(key int) *Node[*BST_Entry[T]] {
	return tree.SearchTree(key, tree.getRoot())
}

func (tree *BinarySearchTree[T]) Insert(data T, key int) {
	entry := BST_Entry[T]{Value: data, Key: key}

	if tree.IsEmpty() {
		nodo := Node[*BST_Entry[T]]{Data: &entry}
		tree.AddRoot(&nodo)
		return
	}

	tree.AddEntry(tree.getRoot(), &entry)
}

func (tree *BinarySearchTree[T]) RemoveByKey(key int) T {

	nodo := tree.Find(key)
	entry := nodo.Data

	if nodo.hasLeft() && nodo.hasRight() {

		precedesor := tree.Predecesor(nodo)
		nodo.Data = precedesor.Data

		tree.Remove(precedesor)

	} else {
		tree.Remove(nodo)
	}

	return entry.GetData()
}

func (tree *BinarySearchTree[T]) SearchTree(key int, nodo *Node[*BST_Entry[T]]) *Node[*BST_Entry[T]] {

	var entry *BST_Entry[T] = nodo.Data

	if key < entry.GetKey() {
		return tree.SearchTree(key, nodo.Left)
	}

	if key > entry.GetKey() {
		return tree.SearchTree(key, nodo.Right)
	}

	return nodo

}

func (tree *BinarySearchTree[T]) AddEntry(nodo *Node[*BST_Entry[T]], entry *BST_Entry[T]) {

	nuevoNodo := Node[*BST_Entry[T]]{Data: entry}

	if entry.GetKey() < nodo.Data.GetKey() {

		if nodo.hasLeft() {
			tree.AddEntry(nodo.Left, entry)
		} else {
			nodo.Left = &nuevoNodo
		}

	} else {
		if nodo.hasRight() {
			tree.AddEntry(nodo.Right, entry)
		} else {
			nodo.Right = &nuevoNodo
		}

	}

}

func (tree *BinarySearchTree[T]) MaxNode(nodo *Node[*BST_Entry[T]]) *Node[*BST_Entry[T]] {
	if nodo.hasRight() {
		return tree.MaxNode(nodo.Right)
	} else {
		return nodo
	}
}

func (tree *BinarySearchTree[T]) Predecesor(nodo *Node[*BST_Entry[T]]) *Node[*BST_Entry[T]] {
	return tree.MaxNode(nodo.Left)
}

func (tree *BinarySearchTree[T]) PrintInorder(nodo *Node[*BST_Entry[T]], callback func(T)) {

	if nodo.hasLeft() {
		tree.PrintInorder(nodo.Left, callback)
	}

	if nodo.hasRight() {
		tree.PrintInorder(nodo.Right, callback)
	}

	callback(nodo.Data.GetData())

}
