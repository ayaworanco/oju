package utils

import (
	"testing"
)

func TestGenerateId(t *testing.T) {
	id := GenerateId()
	if id == "" {
		t.Error("Should not be empty")
	}
}

func TestMapPutBasic(t *testing.T) {
	simple_map := map[string]string{"hello": "world"}

	new_map := MapPut[string, string](simple_map, "name", "john")

	if new_map["name"] != "john" {
		t.Error("MapPut not working")
	}
}
