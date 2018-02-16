package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) Print() {
	fmt.Println(n.Value)
}

func (n *Node) SetValue(val int) {
	if n == nil {
		fmt.Println("Setting Value to a nil node")
		return
	}
	n.Value = val
}

//中序遍历（先左后根再右）
func (n *Node) Traverse() {
	n.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

//函数式中序遍历
func (n *Node) TraverseFunc(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (n *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		n.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}

func CreateNode(val int) *Node {
	return &Node{Value:val}
}
