package main

import (
	"oluwoye/internal/querier"
	"testing"
)

func TestQueries(t *testing.T) {
	tests := []struct {
		description string
		input       string
		valid       bool
	}{
		{
			description: "Valid query",
			input:       "'$ipv4 eq 54.36.149.41 and $status_code eq 400'",
			valid:       true,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.description, func(t *testing.T) {
			result, error := querier.Parse(scenario.input)
			if error != nil {
				t.Error(error.Error())
			}

			if result == nil {
				t.Error("Invalid query")
			}
		})
	}
}
