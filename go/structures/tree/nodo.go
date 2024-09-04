package tree

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
