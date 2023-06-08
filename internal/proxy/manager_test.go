package proxy

import (
	//"fmt"

	"testing"

	"oju/internal/config"
)

func load_config() config.Config {
	config_json := `
{
  "allowed_applications": [
    {
      "name": "worker",
      "app_key": "abc@123",
			"host": "http://worker.api.svc.cluster.local"
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
		Payload: `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`,
	}

	manager.Redirect("abc@123", message)
	traces := manager.GetAppTraces("abc@123")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}
