package entity

import "testing"

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		name string
		cep  string
		want bool
	}{
		{"valid CEP", "12345678", true},
		{"less than 8 digits", "1234567", false},
		{"more than 8 digits", "123456789", false},
		{"contains letter", "1234a678", false},
		{"contains symbol", "1234-678", false},
		{"empty string", "", false},
		{"all zeros", "00000000", true},
		{"all nines", "99999999", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidCEP(tt.cep)
			if got != tt.want {
				t.Errorf("IsValidCEP(%q) = %v; want %v", tt.cep, got, tt.want)
			}
		})
	}
}
