package querier

import (
	"oluwoye/internal/parser"
	"os"
	"strings"
	"testing"
)

func TestValidQuery(t *testing.T) {
	input := "'$ipv4 eq 54.36.149.41 and $status_code eq 200'"
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

	result, error := Parse(input, log_groups)
	if error != nil {
		t.Error(error.Error())
	}

	if !result {
		t.Error("Scenario not valid")
	}
}
