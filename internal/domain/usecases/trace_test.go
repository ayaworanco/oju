package usecases

import "testing"

func TestGenerateTrace(t *testing.T) {
	payload := `{"resource":"app-key-test","action":"action-test","target":"target-test","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

	trace, trace_error := ParseTrace(payload)

	if trace_error != nil {
		t.Error("should not be any error")
	}

	if trace.GetId() == "" {
		t.Error("trace id should not be empty")
	}

}

func TestShouldGenerateErrorWithoutResource(t *testing.T) {
	payload := `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

	trace, trace_error := ParseTrace(payload)

	if trace_error == nil {
		t.Error("Should be errors")
	}

	if trace.Resource != "" {
		t.Error("Resource need to be empty")
	}
}
