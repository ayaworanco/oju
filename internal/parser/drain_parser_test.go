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

func TestDrainParsing(t *testing.T) {
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for _, log := range logs {
		DrainParse(tree, log)
	}
}
