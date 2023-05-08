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
		var operator string

		for range parts {
			length := len(parts)
			if length == 0 {
				break
			}

			operator = parts[1].Data["operator"].(string)

			parts[0].set_result(log_groups)
			parts[2].set_result(log_groups)

			final_result = logical_operation(parts[0].Result, parts[2].Result, operator)

			next_node := query_tree.children[length-1]
			previous_node := query_tree.children[length-2]

			if previous_node.is_operator() && next_node.is_term() {
				next_node.set_result(log_groups)
				final_result = logical_operation(final_result, next_node.Result, previous_node.Data["operator"].(string))
			}
			parts = query_tree.children[3:]
		}
		return final_result
	}
	return false
}
