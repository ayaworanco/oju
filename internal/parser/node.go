package parser

type node struct {
	data     interface{}
	children map[string]*node
}

func new_node(label string, data interface{}) *node {
	return &node{
		data:     data,
		children: make(map[string]*node),
	}
}

func (node *node) add_child(parts []string, log_message string, id int) {
	if len(node.children) > MAX_CHILD {
		return
	}

	if len(parts) == 0 {
		add_log_group(node, log_message, id)
		return
	}

	var path string
	if len(parts) == 1 {
		path = parts[0]
	} else {
		path = parts[:len(parts)-1][0]
	}

	parts = remove_unique_simbol(parts)

	if has_digit(path) {
		path = "*"
	}

	child, ok := node.children[path]
	if !ok {
		child = new_node(path, path)
		node.children[path] = child
	}

	child.add_child(parts[1:], log_message, id)
}
