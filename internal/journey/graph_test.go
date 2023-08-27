package journey

import (
	"oju/internal/config"
	"testing"
)

func TestGraphGenerate(t *testing.T) {
	apps := []config.Application{
		{
			Name:   "test_a",
			AppKey: "test_a",
			Host:   "http://test_a.svc.cluster.local",
		},
	}

	graph := new_graph(apps)

	if len(graph.vertices) != 0 {
		t.Error("vertices are created")
	}
}
