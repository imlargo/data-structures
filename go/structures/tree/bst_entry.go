package tree

type BST_Entry[T any] struct {
	Key   int
	Value T
}

func (entry *BST_Entry[T]) GetKey() int {
	return entry.Key
}

func (entry *BST_Entry[T]) GetData() T {
	return entry.Value
}

func (entry *BST_Entry[T]) SetData(data T) {
	entry.Value = data
}

func (entry *BST_Entry[T]) SetKey(key int) {
	entry.Key = key
}
