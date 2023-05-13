package querier

import (
	"errors"
	"oluwoye/internal/parser"
)

type query_tree struct {
	children []*node
}

func new_query_tree(children []*node) (*query_tree, error) {
	if len(children) == 0 {
		return &query_tree{}, errors.New("nothing parsed")
	}

	return &query_tree{
		children: children,
	}, nil
}

func (query_tree *query_tree) resolve(log_groups []*parser.LogGroup) bool {
	if len(query_tree.children) == 1 {
		query_tree.children[0].set_result(log_groups)
		return query_tree.children[0].Result
	}

	if len(query_tree.children) == 2 {
		return false
	}

	if len(query_tree.children) > 2 {
		parts := query_tree.children[:3]
		var final_result bool
		is_final_result_setted := false
		var operator string

		length := len(parts)
		for range parts {
			if len(parts) == 0 {
				break
			}

			operator = get_operator(parts)
			set_term_result(parts, log_groups)

			if is_final_result_setted {
				next_node := query_tree.children[length-1]
				previous_node := query_tree.children[length-2]

				if previous_node.is_operator() && next_node.is_term() {
					next_node.set_result(log_groups)
					final_result = logical_operation(final_result, next_node.Result, previous_node.Data["operator"].(string))
				}

				if length >= 3 {
					parts = parts[2:]
				} else {
					parts = parts[3:]
				}

			} else {
				final_result = logical_operation(parts[0].Result, parts[2].Result, operator)
				is_final_result_setted = true
				next_node := query_tree.children[length-1]
				previous_node := query_tree.children[length-2]

				if previous_node.is_operator() && next_node.is_term() {
					next_node.set_result(log_groups)
					final_result = logical_operation(final_result, next_node.Result, previous_node.Data["operator"].(string))
				}
				parts = query_tree.children[3:]
			}

		}
		return final_result
	}
	return false
}

func set_term_result(parts []*node, log_groups []*parser.LogGroup) {
	for _, node := range parts {
		if node.is_term() {
			node.set_result(log_groups)
		}
	}
}

func get_operator(parts []*node) string {
	for _, node := range parts {
		if node.is_operator() {
			return node.Data["operator"].(string)
		}
	}
	return ""
}
