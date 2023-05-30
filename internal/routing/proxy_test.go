package routing

import (
	//"fmt"
	"fmt"
	"testing"

	"oju/internal/application"
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
	proxy := NewProxy(config.AllowedApplications)

	message := application.Message{
		Type:    "TRACE",
		Payload: `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`,
	}

	proxy.Redirect("abc@123", message)
	app, app_error := proxy.GetApp("abc@123")

	if app_error != nil {
		t.Fatal(app_error)
	}

	traces := <-app.GetTraces()
	fmt.Printf("%#v", traces)

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}
