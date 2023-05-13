package querier

import (
	"errors"
	"oluwoye/internal/parser"
	"regexp"
	"strings"
)

func Parse(query string, log_groups []*parser.LogGroup) (bool, error) {
	query = strings.Trim(query, "'")
	tokens := strings.Split(query, " ")

	if contains(logical_operators(), tokens[0]) {
		return false, errors.New("first token is logical")
	}

	var nodes []*node

	for range tokens {
		if len(tokens) == 0 {
			break
		}

		if len(tokens) == 1 && contains(logical_operators(), tokens[0]) {
			return false, errors.New("last token is logical")
		}
		parts := tokens[:3]
		var node node
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

	tree, tree_error := new_query_tree(nodes)
	if tree_error != nil {
		return false, tree_error
	}
	return tree.resolve(log_groups), nil

}

func logical_operators() []string {
	return []string{
		"and",
		"or",
		"eq",
		"diff",
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
	case "and":
		return value1.(bool) && value2.(bool)
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

func variables() map[string]regexp.Regexp {
	return map[string]regexp.Regexp{
		"$ipv4":        *regexp.MustCompile("^(?P<ipv4>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"),
		"$status_code": *regexp.MustCompile("^(?P<status_code>[1-5][0-9][0-9])"),
	}
}
