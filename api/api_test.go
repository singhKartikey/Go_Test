package api

import "testing"

func TestUpperCase(t *testing.T) {
	tests := []struct {
		name string
		req  string
	}{
		{
			name: "Happy Path",
			req:  "upper",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := UpperCase(tt.req)
			if str != "UPPER" {
				t.Fatalf("Failed.")
			}
		})
	}
}
