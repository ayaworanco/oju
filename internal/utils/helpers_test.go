package utils

import "testing"

func TestGenerateId(t *testing.T) {
	id := GenerateId()
	if id == "" {
		t.Error("Should not be empty")
	}
}
