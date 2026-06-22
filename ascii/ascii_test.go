package ascii

import (
	"testing"
)

func TestGenerateASCII_ValidInputs(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		banner     string
		wantStatus int
	}{
		{"Standard text request", "Hello World", "standard", 200},
		{"Shadow text request", "$% \"=", "shadow", 200},
		{"Thinkertoy request", "123 T/fs#R", "thinkertoy", 200},
		{"Special characters standard", "{123}\n<Hello> (World)!", "standard", 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, status := GenerateASCII(tt.text, tt.banner)
			if status != tt.wantStatus {
				t.Errorf("GenerateASCII() status = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}

func TestGenerateASCII_InvalidInputs(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		banner     string
		wantStatus int
	}{
		{"Non-ASCII Unicode Character", "Hello, 世界", "standard", 400},
		{"Missing Template File Error", "Valid Text", "nonexistent-banner", 404},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, status := GenerateASCII(tt.text, tt.banner)
			if status != tt.wantStatus {
				t.Errorf("GenerateASCII() status = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}