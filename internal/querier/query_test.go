package querier

import (
	"oluwoye/internal/parser"
	"os"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		description string
		input       string
		is_valid    bool
	}{
		{
			description: "Invalid query starting with logical operator",
			input:       "'and $ipv4 eq 54.36.149.41'",
			is_valid:    false,
		},
		{
			description: "Invalid query with logical operator in the end",
			input:       "'$ipv4 eq 54.36.149.41 and'",
			is_valid:    false,
		},
		{
			description: "Valid query with one term",
			input:       "'$ipv4 eq 54.36.149.41'",
			is_valid:    true,
		},
		{
			description: "Valid query with two terms",
			input:       "'$ipv4 eq 54.36.149.41 and $status_code eq 200'",
			is_valid:    true,
		},
		{
			description: "Valid query with 2 or more terms",
			input:       "'$ipv4 eq 54.36.149.41 and $status_code eq 200 and $status_code diff 500'",
			is_valid:    true,
		},
	}

	tree := parser.NewTree(10)

	file, file_error := os.ReadFile("testdata/query_test.log")
	if file_error != nil {
		t.Error(file_error.Error())
	}

	log := string(file)
	logs := strings.Split(log, "\n")

	for id, registry := range logs {
		parser.ParseLog(tree, registry, id)
	}

	log_groups := tree.GetLogGroups(tree.GetRoot())

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result, error := Parse(test.input, log_groups)

			if error != nil {
				t.Error(error.Error())
			}

			if test.is_valid && result {
				t.Error("Scenario not valid")
			}
		})
	}
}
