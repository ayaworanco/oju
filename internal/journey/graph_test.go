package journey

import (
	"oju/internal/tracer"
	"testing"
)

func TestUpdateGraph(t *testing.T) {
	graph := NewGraph(make(map[string]Vertex))

	trace := tracer.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	command := NewInsertActionCommand(trace)

	graph = UpdateGraph(graph, command)

	if len(graph.Vertices) == 0 {
		t.Error("vertices length should be greatet than 0")
	}
}
