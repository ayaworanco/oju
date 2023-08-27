package journey

import (
	"testing"
	"oju/internal/config"
	"oju/internal/tracer"
)

type test_suite struct {
	apps []config.Application
	system system
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
		apps: apps,
		system: NewSystem(apps),
	}

}

func TestInsertAction(t *testing.T) {
	suite := setup()

	trace := tracer.Trace{
		AppKey: "test1",
		Name: "bhaskara",
		Service: "delta",
		Attributes: make(map[string]string),
	}

	suite.system.InsertAction(trace)
}
