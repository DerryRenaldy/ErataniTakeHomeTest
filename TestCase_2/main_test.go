package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"kodok", "kodok", true},
		{"kodokkodok", "kodokkodok", true},
		{"jinggaaku", "jinggaaku", false},
		{"123321", "123321", true},
		{"Aba", "Aba", true},
		{"racecar", "racecar", true},
		{"level", "level", true},
		{"rotor", "rotor", true},
		{"madam", "madam", true},
		{"hello", "hello", false},
		{"world", "world", false},
		{"123", "123", false},
		{"test", "test", false},
		{"CaseInsensitive: Aba", "aba", true},
		{"CaseInsensitive: RADAR", "radar", true},
		{"Single character: a", "a", true},
		{"Empty string", "", true},
		{"Single character: z", "z", true},
		{"Single character: 5", "5", true},
		{"Single character: B", "B", true},
		{"Two characters same: aa", "aa", true},
		{"Two characters different: ab", "ab", false},
		{"Two characters same: 11", "11", true},
		{"Two characters different: 12", "12", false},
		{"Three characters: aba", "aba", true},
		{"Three characters: abc", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
