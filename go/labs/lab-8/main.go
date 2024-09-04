package main

import "data-structures/structures/tree"

func main() {
	var root tree.Node[int] = tree.Node[int]{Data: 1}
	arbol := tree.BinarySearchTree[int]{
		Root: &root,
	}

	println("Root: ", arbol.Root.Data)

}
