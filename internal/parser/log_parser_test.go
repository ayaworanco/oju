package parser

import (
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
	tree := NewTree(10)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for id, log := range logs {
		ParseLog(tree, log, id)
	}

	log_group := tree.GetLogGroups(tree.root)
	if len(log_group) != 3 {
		t.Error("Should be a length 3")
	}

}
