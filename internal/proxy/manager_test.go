package proxy

import (
	"testing"

	"oju/internal/config"
	"oju/internal/tracer"
)

func load_config() config.Config {
	config_json := `
{
  "allowed_applications": [
    {
      "name": "bhaskara",
      "app_key": "abc@123",
			"host": "http://bhaskara.api.svc.cluster.local"
    },
	{
		"name": "delta",
		"app_key": "def@321",
		"host": "http://delta.api.svc.cluster.local"
	}
  ]
}
`

	config, _ := config.BuildConfig([]byte(config_json))
	return config
}

func TestUpAllAllowedApplicationsByProxy(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: `{"app_key": "abc@123","name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`,
	}

	manager.Redirect("abc@123", message)
	traces := manager.GetAppTraces("abc@123")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}

func TestGettingService(t *testing.T) {
	metadata := Metadata{Key: "delta", Host: "http://delta.api.svc.cluster.local"}

	trace_service := tracer.Trace{
		AppKey: "abc@123",
		Name: "calculate-delta",
		Service: metadata.Key,
		Attributes: make(map[string]string, 0),
	}

	trace_attributes := tracer.Trace{
		AppKey: "abc@123",
		Name: "calculate-delta",
		Service: "",
		Attributes: map[string]string{
			"http.url": metadata.Host,
		},
	}

	service_filled, service_error_filled := get_service(trace_service, []Metadata{metadata})

	if service_error_filled != nil {
		t.Error("Should be a service name")
	}

	if service_filled == "" {
		t.Error("Should be filled up")
	}

	service_attributes, service_error_attributes := get_service(trace_attributes, []Metadata{metadata})
	if service_error_attributes != nil {
		t.Error("Should be a service name")
	}

	if service_attributes == "" {
		t.Error("Should be filled up")
	}
}

func TestTwoTracesByDifferentServicesByAppKeyField(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: `{"app_key":"abc@123","name":"calculate-delta","service":"def@321","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`,
	}

	message_delta := ApplicationMessage{
		Type:    "TRACE",
		Payload: `{"app_key":"def@321","name":"check-delta","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`,
	}

	manager.Redirect("abc@123", message)
	manager.Redirect("def@321", message_delta)
	traces := manager.GetAppTraces("abc@123")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}

	var bhaskara_trace *tracer.Trace

	for _, trace := range traces {
		bhaskara_trace = trace
		break
	}

	children := bhaskara_trace.GetChildren()

	if len(children) == 0 {
		t.Fatal("this children must be filled")
	}
}
