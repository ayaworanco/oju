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
      "app_key": "bhaskara",
			"host": "http://bhaskara.api.svc.cluster.local"
    },
	{
		"name": "delta",
		"app_key": "delta",
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
		Payload: `{"app_key": "bhaskara","name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`,
	}

	manager.Redirect("bhaskara", message)
	traces := manager.GetAppTraces("bhaskara")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}

func TestTwoTracesByDifferentServicesByAppKeyField(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: `{"app_key":"bhaskara","name":"calculate-delta","service":"delta","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`,
	}

	message_delta := ApplicationMessage{
		Type:    "TRACE",
		Payload: `{"app_key":"delta","name":"check-delta","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`,
	}

	manager.Redirect("bhaskara", message)
	manager.Redirect("delta", message_delta)
	traces := manager.GetAppTraces("bhaskara")

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
