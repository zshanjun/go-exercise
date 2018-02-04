package main

import (
	"zshanjun/go-exercise/tree"
	"fmt"
	"zshanjun/go-exercise/tree/mytree"
)

func main() {
	root := tree.CreateNode(1)
	root.Left = &tree.Node{Value:2}
	root.Right = &tree.Node{Value:3}
	root.Left.Left = new(tree.Node)
	root.Left.Left.Value = 5
	root.Right.Left = tree.CreateNode(4)
	root.Traverse()
	fmt.Println()

	n := mytree.MyNode{root}
	n.PostOrder()
	fmt.Println()
}
