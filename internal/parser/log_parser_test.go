package parser

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
Command has run successfully
`

const TEST_LOG_MESSAGE = "54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

func TestLogParsing(t *testing.T) {
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for id, log := range logs {
		ParseLog(tree, log, id)
	}

	log_group := tree.root.children["3"].children["Temperature"].children["*"].children["exceeds"].children["log_group_3"].data.(*LogGroup)

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
	for id, log := range logs {
		ParseLog(tree, log, id)
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
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for id, log := range logs {
		ParseLog(tree, log, id)
	}

	log_group := tree.GetLogGroups(tree.root)
	fmt.Printf("%#v", log_group)

}
