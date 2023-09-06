package usecases

import "testing"

func TestReturnAllowedApps(t *testing.T) {
	config_json := `	
{
	"resources": [
	{
	"name": "worker",
	"key": "37F129EF-687E-47AC-B0C4-CBF6516B37ED",
	"host": "http://localhost"
	}
	]
}
	`

	config, load_error := BuildConfig([]byte(config_json))
	if load_error != nil {
		t.Error(load_error)
	}

	if len(config.Resources) == 0 {
		t.Error("Should be at least 1")
	}
}
