package usecases

import (
	"oju/internal/domain/entities"
	"testing"
)

func TestUpdateGraph(t *testing.T) {
	graph := new_graph(make(map[string]entities.Vertex))

	trace := entities.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	command := entities.NewInsertActionCommand(trace)
	graph = update_graph(graph, command)

	if len(graph.Vertices) == 0 {
		t.Error("vertices len should be greater than 0")
	}
}
