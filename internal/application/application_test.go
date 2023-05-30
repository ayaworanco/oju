package application

import (
	"testing"
)

func TestStartProcessAndSendLog(t *testing.T) {
	app := Start(10, Metadata{
		Key:        "abc@123",
		Host:       "http://app.api.svc.cluster.local",
		WatchQuery: "",
	})

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
	groups := app.tree.GetLogGroups(app.tree.GetRoot())

	if len(groups) == 0 {
		t.Error("Should have been parsed")
	}
}

func TestStartProcessAndSendTrace(t *testing.T) {
	app := Start(10, Metadata{
		Key:        "abc@123",
		Host:       "http://app.api.svc.cluster.local",
		WatchQuery: "",
	})

	message := Message{
		Type:    "TRACE",
		Payload: `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`,
	}

	app.SendMessage(message)
}
