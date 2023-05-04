package querier

import (
	"regexp"
	"strings"
)

type QueryTree struct {
	Query    string
	Children []*Node
}

type Node struct {
	Data map[string]interface{}
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
		Children: nodes,
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

func variables() map[string]regexp.Regexp {
	return map[string]regexp.Regexp{
		"$ipv4": *regexp.MustCompile("^(?P<ipv4>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"),
	}

}
