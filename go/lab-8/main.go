package main

import "lab-8/tree"

func main() {
	var root tree.Node[int] = tree.Node[int]{Data: 1}
	arbol := tree.BinaryTree[int]{
		Root: &root,
	}

	println("Root: ", arbol.Root.Data)

}
