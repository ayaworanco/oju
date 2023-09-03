package armazen

import (
	"oju/internal/config"
	"oju/internal/journey"
	"oju/internal/system"
	"oju/internal/tracer"
	"os"
	"testing"
)

type test_suite struct {
	apps   []config.Application
	system system.System
}

func setup() test_suite {
	apps := []config.Application{
		{
			Name:   "test_a",
			AppKey: "test_a",
			Host:   "http://test_a.svc.cluster.local",
		},
	}

	return test_suite{
		apps:   apps,
		system: system.NewSystem(apps),
	}

}

func TestEncodingSystemToBytes(t *testing.T) {
	suite := setup()

	trace := tracer.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := journey.NewInsertActionCommand(trace)
	system.Send(suite.system, message)

	system_bytes, encoding_error := EncodeSystemToBytes(suite.system)
	if encoding_error != nil {
		t.Error(encoding_error.Error())
	}

	if len(system_bytes) == 0 {
		t.Error("length of system bytes should be greater than 0")
	}
}

func TestCompressingSystemToBytes(t *testing.T) {
	suite := setup()

	trace := tracer.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := journey.NewInsertActionCommand(trace)
	system.Send(suite.system, message)

	system_bytes, encoding_error := EncodeSystemToBytes(suite.system)
	if encoding_error != nil {
		t.Error(encoding_error.Error())
	}

	if len(system_bytes) == 0 {
		t.Error("length of system bytes should be greater than 0")
	}

	zipped_bytes := Compress(system_bytes)
	if len(zipped_bytes) == 0 {
		t.Error("length of zipped bytes should be greater than 0")
	}
}

func TestSavingSystemOnDisk(t *testing.T) {
	t.Setenv("OUTPUT_DATA_DIR", "/tmp/")
	suite := setup()

	trace := tracer.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := journey.NewInsertActionCommand(trace)
	system.Send(suite.system, message)

	system_bytes, encoding_error := EncodeSystemToBytes(suite.system)
	if encoding_error != nil {
		t.Error(encoding_error.Error())
	}

	if len(system_bytes) == 0 {
		t.Error("length of system bytes should be greater than 0")
	}

	zipped_bytes := Compress(system_bytes)
	if len(zipped_bytes) == 0 {
		t.Error("length of zipped bytes should be greater than 0")
	}

	write_error := WriteToFile(zipped_bytes)
	if write_error != nil {
		t.Error(write_error.Error())
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
