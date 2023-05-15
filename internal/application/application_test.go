package application

import (
	"testing"

	"oluwoye/internal/parser"
)

func TestStartProcessAndSendLog(t *testing.T) {
	tree := parser.NewTree(8)
	app := StartApplication(tree, "abc@123", "")

	message := Message{
		Type:    "LOG",
		Payload: "Temperature (43C) exceeds",
	}

	app.SendMessage(message)
	groups := tree.GetLogGroups(tree.GetRoot())

	if len(groups) == 0 {
		t.Error("Should have been parsed")
	}
}
