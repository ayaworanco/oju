package journey

import (
	"testing"
	"oju/internal/tracer"
)

func TestUpdateGraph(t *testing.T) {
	graph := new_graph(make(map[string]vertex))

	trace := tracer.Trace{
		AppKey: "test1",
		Name: "bhaskara",
		Service: "delta",
		Attributes: make(map[string]string),
	}

	command := InsertActionCommand{
		Type: INSERT_ACTION,
		Data: trace,
	}

	graph = update_graph(graph, command)

	if len(graph.vertices) == 0 {
		t.Error("vertices length should be greatet than 0")
	}
}
