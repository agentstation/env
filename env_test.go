package env

import (
	"testing"
)

func TestEnvironmentString(t *testing.T) {
	tests := []struct {
		env      Environment
		expected string
	}{
		{PROD, "Production"},
		{STAGE, "Staging"},
		{DEV, "Development"},
		{"UNKNOWN", "Unknown"},
		{"", "Unknown"}, // Edge case: empty string
	}

	for _, test := range tests {
		result := test.env.String()
		if result != test.expected {
			t.Errorf("For env %v, expected %s, got %s", test.env, test.expected, result)
		}
	}
}

func TestIsDev(t *testing.T) {
	tests := []struct {
		env      Environment
		expected bool
	}{
		{DEV, true},
		{PROD, false},
		{STAGE, false},
		{"UNKNOWN", false},
		{"", false}, // Edge case: empty string
	}

	for _, test := range tests {
		result := test.env.IsDev()
		if result != test.expected {
			t.Errorf("For env %v, expected IsDev to be %v, got %v", test.env, test.expected, result)
		}
	}
}

func TestIsStage(t *testing.T) {
	tests := []struct {
		env      Environment
		expected bool
	}{
		{STAGE, true},
		{DEV, false},
		{PROD, false},
		{"UNKNOWN", false},
		{"", false}, // Edge case: empty string
	}

	for _, test := range tests {
		result := test.env.IsStage()
		if result != test.expected {
			t.Errorf("For env %v, expected IsStage to be %v, got %v", test.env, test.expected, result)
		}
	}
}

func TestIsProd(t *testing.T) {
	tests := []struct {
		env      Environment
		expected bool
	}{
		{PROD, true},
		{STAGE, false},
		{DEV, false},
		{"UNKNOWN", false},
		{"", false}, // Edge case: empty string
	}

	for _, test := range tests {
		result := test.env.IsProd()
		if result != test.expected {
			t.Errorf("For env %v, expected IsProd to be %v, got %v", test.env, test.expected, result)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected Environment
	}{
		{"PROD", PROD},
		{"production", PROD},
		{"STAGE", STAGE},
		{"staging", STAGE},
		{"DEV", DEV},
		{"development", DEV},
		{"invalid", PROD}, // default case
		{"", PROD},        // Edge case: empty string
	}

	for _, test := range tests {
		result := Parse(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestSetAndCurrent(t *testing.T) {
	environments := []Environment{DEV, STAGE, PROD}

	for _, env := range environments {
		Set(env)
		if Current() != env {
			t.Errorf("Expected current environment to be %v, got %v", env, Current())
		}
	}

	// Edge case: setting an unknown environment
	Set("UNKNOWN")
	if Current() != "UNKNOWN" {
		t.Errorf("Expected current environment to be UNKNOWN, got %v", Current())
	}
}
