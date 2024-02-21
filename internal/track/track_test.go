package track

import "testing"

func TestGenerateTrack(t *testing.T) {
	payload := `{"resource":{"name":"track-name","action":"track-action"},"target":{"name":"target-name","action":"target-action"},"attributes":{"http.method":"POST"}}`
	track, track_error := Parse(payload)

	if track_error != nil {
		t.Error("should not be any error")
	}

	if track.GetID() == "" {
		t.Error("track id should not be empty")
	}
}

func TestShouldGenerateErrorWithoutResource(t *testing.T) {
	payload := `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`
	track, track_error := Parse(payload)

	if track_error == nil {
		t.Error("should be errors")
	}

	if track.Resource.Name != "" {
		t.Error("resource should be empty")
	}
}
