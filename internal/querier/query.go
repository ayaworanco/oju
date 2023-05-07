package querier

import (
	"regexp"
	"strings"
)

type QueryTree struct {
	Query    string
	children []*Node
}

type Node struct {
	Data   map[string]interface{}
	Result bool
}

func variables() map[string]regexp.Regexp {
	return map[string]regexp.Regexp{
		"$ipv4":        *regexp.MustCompile("^(?P<ipv4>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"),
		"$status_code": *regexp.MustCompile("^(?P<status_code>[1-5][0-9][0-9])"),
	}
}

func (node *Node) set_result(message string) {
	regex, ok_variable := node.Data["variable"]
	operator, ok_operator := node.Data["operator"]
	value, ok_value := node.Data["value"]

	var variable_result string

	if ok_variable && ok_value && ok_operator {
		expression := regex.(regexp.Regexp)
		if expression.MatchString(message) {
			variable_result = expression.FindString(message)
		} else {
			node.Result = false
		}
		node.Result = logical_operation(variable_result, value, operator.(string))
	} else {
		node.Result = false
	}

}

func (query_tree *QueryTree) Resolve(message string) bool {

	if len(query_tree.children) == 1 {
		query_tree.children[0].set_result(message)
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
			if len(parts) == 0 {
				break
			}
			first_term := parts[0]
			operator = parts[1].Data["operator"].(string)
			second_term := parts[2]

			// Set results
			first_term.set_result(message)
			second_term.set_result(message)

			final_result = logical_operation(first_term.Result, second_term.Result, operator)

			next := query_tree.children[len(parts)-1]
			previous := query_tree.children[len(parts)-2]

			if is_operator(previous) && is_term(next) {
				next.set_result(message)
				final_result = logical_operation(final_result, next.Result, previous.Data["operator"].(string))
			}
			parts = query_tree.children[3:]
		}
		return final_result
	}
	return false
}

func is_term(node *Node) bool {
	_, ok_variable := node.Data["variable"]
	_, ok_operator := node.Data["operator"]
	_, ok_value := node.Data["value"]
	return ok_variable && ok_operator && ok_value
}

func is_operator(node *Node) bool {
	_, ok_variable := node.Data["variable"]
	_, ok_operator := node.Data["operator"]
	_, ok_value := node.Data["value"]
	return !ok_variable && ok_operator && !ok_value
}

func Parse(query string) (*QueryTree, error) {
	query = strings.Trim(query, "'")
	tokens := strings.Split(query, " ")

	var nodes []*Node

	for i := 0; i <= len(tokens); i++ {
		parts := tokens[:3]
		var node Node
		m := make(map[string]interface{})

		if contains(logical_operators(), parts[0]) {
			m["operator"] = parts[0]
			node.Data = m
			nodes = append(nodes, &node)
			new_tokens := tokens[1:]
			tokens = new_tokens[:len(tokens)-1]
			continue
		}
		for _, part := range parts {
			if strings.HasPrefix(part, "$") {
				m["variable"] = variables()[part]
			} else if contains(logical_operators(), part) {
				m["operator"] = part
			} else {
				m["value"] = part
			}
		}
		node.Data = m
		nodes = append(nodes, &node)
		tokens = tokens[3:]
	}

	return &QueryTree{
		Query:    query,
		children: nodes,
	}, nil
}

func logical_operators() []string {
	return []string{
		"and",
		"or",
		"eq",
	}
}

func contains(elements []string, value string) bool {
	for _, element := range elements {
		if value == element {
			return true
		}
	}
	return false
}

func logical_operation(value1, value2 interface{}, operator string) bool {
	switch operator {
	case "eq":
		return value1 == value2
	case "diff":
		return value1 != value2
	case "gt":
		return value1.(int) > value2.(int)
	case "lt":
		return value1.(int) < value2.(int)
	default:
		return false
	}
}
