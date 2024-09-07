package tree

import (
	"math"
	"strings"
)

type Identificador[T any] struct {
	nodo  *Node[*BST_Entry[T]]
	level int
	x     int
	align string
}

func (tree *BinarySearchTree[T]) PrintTree(getValueCallback func(T) string) {
	root := tree.getRoot()

	nlevels := tree.Height(root)
	width := int(math.Pow(2, float64(nlevels+1)))

	q := []Identificador[T]{{root, 0, width, "c"}}

	levels := make([][]Identificador[T], 0)

	for len(q) > 0 {

		// pop
		current := q[0]
		q = q[1:]

		if current.nodo != nil {
			if len(levels) <= current.level {
				levels = append(levels, []Identificador[T]{})
			}

			levels[current.level] = append(levels[current.level], Identificador[T]{nodo: current.nodo, level: current.level, x: current.x, align: current.align})

			seg := int(float64(width) / math.Pow(2, float64(current.level+1)))

			q = append(q, Identificador[T]{current.nodo.Left, current.level + 1, current.x - seg, "l"})
			q = append(q, Identificador[T]{current.nodo.Right, current.level + 1, current.x + seg, "r"})

		}
	}

	for i, level := range levels {
		pre := 0
		preline := 0
		linestr := ""
		pstr := ""
		seg := int(float64(width) / math.Pow(2.0, float64(i+1)))

		for _, element := range level {
			valstr := getValueCallback(element.nodo.Data.GetData())

			if element.align == "r" {
				a := strings.Repeat(" ", element.x-preline-1-seg-int(seg/2))
				b := strings.Repeat("¯", seg+int(seg/2))
				linestr += a + b + "\\"
				preline = element.x

			} else if element.align == "l" {

				a := strings.Repeat(" ", element.x-preline-1)
				b := strings.Repeat("¯", seg+int(seg/2))
				linestr += a + "/" + b

				preline = element.x + seg + int(seg/2)
			}

			max := int(math.Max(float64(element.x-pre-len(valstr)), float64(-1*(element.x-pre-len(valstr)))))
			pstr += strings.Repeat(" ", max) + valstr
			pre = element.x
		}

		println(linestr)
		println(pstr)
	}

}
