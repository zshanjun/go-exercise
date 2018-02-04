package mytree

import "zshanjun/go-exercise/tree"

type MyNode struct {
	Node *tree.Node
}

//后序遍历
func (m *MyNode) PostOrder() {
	if m == nil || m.Node == nil {
		return
	}

	left := MyNode{m.Node.Left}
	left.PostOrder()
	right := MyNode{m.Node.Right}
	right.PostOrder()
	m.Node.Print()
}
