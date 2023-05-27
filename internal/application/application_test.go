package application

import (
	"testing"

	"oju/internal/parser"
)

func TestStartProcessAndSendLog(t *testing.T) {
	tree := parser.NewTree(8)
	app := Start(tree, "abc@123", "")

	messages := []Message{
		{
			Type:    "LOG",
			Payload: "Temperature (43C) exceeds",
		},
		{
			Type:    "LOG",
			Payload: "Temperature (45C) exceeds",
		},
		{
			Type:    "LOG",
			Payload: "Temperature (62C) exceeds",
		},
	}

	for _, message := range messages {
		app.SendMessage(message)
	}
	groups := tree.GetLogGroups(tree.GetRoot())

	if len(groups) == 0 {
		t.Error("Should have been parsed")
	}
}
