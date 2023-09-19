package usecases

import (
	"oju/internal/domain/entities"
	"os"
	"testing"
)

func TestEncodingSystemToBytes(t *testing.T) {
	t.Setenv("OUTPUT_DATA_DIR", "/tmp")
	suite := setup()

	trace := entities.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := entities.NewInsertActionCommand(trace)
	Send(suite.system, message)

	save_error := save(suite.system)
	if save_error != nil {
		t.Error(save_error.Error())
	}
	env_data_dir := os.Getenv("OUTPUT_DATA_DIR") + "/data"

	dir, dir_error := os.ReadDir(env_data_dir)
	if dir_error != nil {
		t.Error(dir_error.Error())
	}

	if len(dir) == 0 {
		t.Error("should have at least one file")
	}
}

func TestLoadSavedSystem(t *testing.T) {
	t.Setenv("OUTPUT_DATA_DIR", "/tmp")
	suite := setup()

	trace := entities.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := entities.NewInsertActionCommand(trace)
	Send(suite.system, message)

	save_error := save(suite.system)
	if save_error != nil {
		t.Error(save_error.Error())
	}
	// env_data_dir := os.Getenv("OUTPUT_DATA_DIR") + "/data"
	sys, error_loading := LoadSystem()
	if error_loading != nil {
		t.Error(error_loading.Error())
	}

	if len(sys.Resources) == 0 {
		t.Error("should at least have 1 resource")
	}

	if len(sys.Graph.Vertices) == 0 {
		t.Error("should at least have 1 graph vertice")
	}
}
