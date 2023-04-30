package parser

import (
	"fmt"
	"strings"
)

type Node struct {
	Data     interface{}
	Children map[interface{}]*Node
}

type Tree struct {
	Root  *Node
	Depth int
}

func NewTree(depth int) *Tree {
	return &Tree{
		Root: &Node{
			Data:     "root",
			Children: make(map[interface{}]*Node, 0),
		},
		Depth: depth,
	}
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:     data,
		Children: make(map[interface{}]*Node),
	}
}

func (tree *Tree) AddOrUpdateLengthLayer(log string) {
	parts := strings.SplitN(log, " ", tree.Depth)
	length := len(parts)
	// tree.Root.Add(path)
	// length_node := &Node{
	// 	Data:     length,
	// 	Children: make([]*Node, 0),
	// }

	child, ok := tree.Root.Children[length]
	if !ok {
		child = NewNode(length)
		// child.Add(parts)
		tree.Root.Children[length] = child

	}
	child.Add(parts[1:])
}

func (node *Node) Add(parts []string) {
	if len(parts) == 0 {
		return
	}

	var path string
	if len(parts) == 1 {
		path = parts[0]
	} else {
		path = parts[:len(parts)-1][0]
	}

	label := path[0]
	child, ok := node.Children[label]
	if !ok {
		child = NewNode(label)
		node.Children[label] = child
	}

	child.Add(parts[1:])
}

func DrainParse(tree *Tree, log string) {
	// this tree needs to be updated every time
	tree.AddOrUpdateLengthLayer(log)
	fmt.Printf("%#v", log)
}
