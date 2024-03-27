package drain

import (
	"fmt"
	"strings"
	"testing"
)

const TEST_LOG_ONE = `
Temperature (41C) exceeds
`

const TEST_LOG_EASY = `Temperature (41C) exceeds
Temperature (43C) exceeds
Command has run successfully`

const TEST_LOG_DIFFICULT = `233.223.117.90 - - [27/Dec/2037:12:00:00 +0530] "DELETE /usr/admin HTTP/1.0" 502 4963 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4380.0 Safari/537.36 Edg/89.0.759.0" 45
162.253.4.179 - - [27/Dec/2037:12:00:00 +0530] "GET /usr/admin/developer HTTP/1.0" 200 5041 "http://www.parker-miller.org/tag/list/list/privacy/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36" 3885
252.156.232.172 - - [27/Dec/2037:12:00:00 +0530] "POST /usr/register HTTP/1.0" 404 5028 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 OPR/73.0.3856.329" 3350`

func TestLogParsing(t *testing.T) {
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for _, log := range logs {
		ParseLog(tree, log)
	}

	log_group := tree.root.children["3"].children["Temperature"].children["*"].children["exceeds"].children["log_group_3"].data.(*LogGroup)
	fmt.Printf("%#v", log_group)

	if log_group.LogEvent != "Temperature * exceeds" {
		t.Error("This log should have wildcards")
	}

	if len(log_group.LogParameters) != 2 {
		t.Error("This log need to have at least 2 parameters")
	}
}

func TestLogParsingWithOneLog(t *testing.T) {
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_ONE, "\n")
	for _, log := range logs {
		ParseLog(tree, log)
	}

	log_group := tree.root.children["3"].children["Temperature"].children["*"].children["exceeds"].children["log_group_3"].data.(*LogGroup)

	if log_group.LogEvent != "Temperature * exceeds" {
		t.Error("This log should have wildcards")
	}

	if len(log_group.LogParameters) != 1 {
		t.Error("This log need to have at least 1 parameters")
	}
}

func TestLogMessageGroup(t *testing.T) {
	tree := NewTree(10)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for _, log := range logs {
		ParseLog(tree, log)
	}

	log_group := tree.GetLogGroups(tree.root)
	if len(log_group) != 2 {
		t.Error("Should be a length 2")
	}
}
