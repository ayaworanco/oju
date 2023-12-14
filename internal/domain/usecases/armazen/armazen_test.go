package armazen

import (
	"fmt"
	"oju/internal/domain/entities"
	"os"
	"testing"
)

func graph_example() map[string][]string {
	return map[string][]string{
		"resource@action_1": []string{"resource@action_2", "resource@action_3"},
		"resource@action_2": []string{"resource@action_4"},
		"resource@action_4": []string{},
	}
}

func TestEncodingGraphToBytes(t *testing.T) {
	t.Setenv("OUTPUT_DATA_DIR", "/tmp/oju")

	graph_example := graph_example()
	save_error := save(graph_example)

	if save_error != nil {
		t.Error(save_error.Error())
	}

	env_data_dir := os.Getenv("OUTPUT_DATA_DIR") + "/data"

	dir_files, dir_load_error := os.ReadDir(env_data_dir)
	if dir_load_error != nil {
		t.Error(dir_load_error.Error())
	}

	if len(dir_files) == 0 {
		t.Error("should have at least one file")
	}

	os.RemoveAll(os.Getenv("OUTPUT_DATA_DIR") + "/data")
}

func TestLoadGraph(t *testing.T) {
	t.Setenv("OUTPUT_DATA_DIR", "/tmp")

	graph_example := graph_example()
	save_error := save(graph_example)

	graph_loaded, error_to_load_graph := Load()
	if error_to_load_graph != nil {
		t.Error(error_to_load_graph.Error())
	}

	if len(graph_loaded) == 0 {
		t.Error("should at leas have 1 vertice")
	}
}
