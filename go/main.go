package main

import "data-structures/structures/tree"

func main() {
	var arbol tree.BinaryTree[int]
	arbol.AddRoot(&tree.Node[int]{Data: 1})

	println("Root: ", arbol.Size)

}
