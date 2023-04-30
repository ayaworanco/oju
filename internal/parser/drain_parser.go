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

type LogGroup struct {
	LogEvent      string
	LogParameters map[int]string
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

func (tree *Tree) AddOrUpdateLengthLayer(log string, id int) {
	parts := strings.SplitN(log, " ", tree.Depth)
	length := len(parts)

	child, ok := tree.Root.Children[length]
	if !ok {
		child = NewNode(length)
		tree.Root.Children[length] = child

	}
	child.Add(parts[1:], log, id)
}

func (node *Node) Add(parts []string, log_message string, id int) {
	if len(parts) == 0 {
		log_group := LogGroup{
			LogEvent:      log_message,
			LogParameters: map[int]string{},
		}
		child, ok := node.Children[log_group]
		if !ok {
			child = NewNode(log_group)
			node.Children[log_group] = child
		}
		// TODO: need to run the similarity function to check if this group is suitable
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

	child.Add(parts[1:], log_message, id)
}

func DrainParse(tree *Tree, log string, id int) {
	// this tree needs to be updated every time
	tree.AddOrUpdateLengthLayer(log, id)
	fmt.Printf("%#v", log)
}
