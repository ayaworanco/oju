package journey

import (
)

/*
type test_suite[T tracer.Trace | string] struct {
	apps []config.Application
	system system[T]
}

func setup[T tracer.Trace | string]() test_suite[T] {
	apps := []config.Application{
		{
			Name:   "test_a",
			AppKey: "test_a",
			Host:   "http://test_a.svc.cluster.local",
		},
	}

	return test_suite[T]{
		apps: apps,
		system: NewSystem[T](apps),
	}

}

func TestInsertAction(t *testing.T) {
	setup[tracer.Trace]()

	trace := tracer.Trace{
		AppKey: "test1",
		Name: "bhaskara",
		Service: "delta",
		Attributes: make(map[string]string),
	}

	message := SystemMessage[tracer.Trace]{
		Type: INSERT_ACTION,
		Data: trace,
	}

	Send[tracer.Trace](message)
}

*/
