package parser

import (
	"fmt"
	"strings"
)

type Tree struct {
	root  *node
	depth int
}

func NewTree(depth int) *Tree {
	return &Tree{
		root: &node{
			data:     "root",
			children: make(map[string]*node, 0),
		},
		depth: depth - 2,
	}
}

func (tree *Tree) GetRoot() *node {
	return tree.root
}

func (tree *Tree) add_or_update_length_layer(log string, id int) {
	parts := strings.SplitN(log, " ", tree.depth)
	length := len(parts)
	label := fmt.Sprint(length)

	child, ok := tree.root.children[label]
	if !ok {
		child = new_node(label, length)
		tree.root.children[label] = child

	}
	child.add_child(parts, log, id)
}

func (tree *Tree) GetLogGroups(node *node) []*LogGroup {
	var log_groups []*LogGroup

	group, is_log_group := node.data.(*LogGroup)

	if is_log_group {
		log_groups = append(log_groups, group)
	} else {
		for _, child := range node.children {
			log_groups = append(log_groups, tree.GetLogGroups(child)[0])
		}
	}

	return log_groups
}
