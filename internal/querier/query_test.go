package querier

import (
	"oluwoye/internal/parser"
	"os"
	"strings"
	"testing"
)

const LOG_MESSAGE = "54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

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
