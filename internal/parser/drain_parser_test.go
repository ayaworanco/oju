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

func TestDrainParsing(t *testing.T) {
	tree := NewTree(8)

	logs := strings.Split(TEST_LOG_EASY, "\n")
	for id, log := range logs {
		DrainParse(tree, log, id)
	}

	log_group := tree.Root.Children["3"].Children["Temperature"].Children["*"].Children["exceeds"].Children["log_group_3"].Data

	fmt.Printf("%#v", log_group)
}
