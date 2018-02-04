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
	if n == nil {
		return
	}
	n.Left.Traverse()
	n.Print()
	n.Right.Traverse()
}

func CreateNode(val int) *Node {
	return &Node{Value:val}
}
