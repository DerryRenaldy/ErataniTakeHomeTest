package main

import (
	"testing"
)

func TestFindKthAccepted(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"k=0", 0, 0},
		{"k=1", 1, 1},
		{"k=2", 2, 2},
		{"k=3", 3, 4},
		{"k=4", 4, 5},
		{"k=5", 5, 7},
		{"k=6", 6, 8},
		{"k=7", 7, 10},
		{"k=8", 8, 11},
		{"k=9", 9, 14},
		{"k=10", 10, 16},
		{"k=20", 20, 32},
		{"k=50", 50, 82},
		{"k=100", 100, 166},
		{"k=500", 500, 832},
		{"k=1000", 1000, 1666},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindAcceptedValue(tt.input)
			if result != tt.expected {
				t.Errorf("FindKthAccepted(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsAccepted(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"1 is accepted", 1, true},
		{"2 is accepted", 2, true},
		{"3 is not (divisible by 3)", 3, false},
		{"4 is accepted", 4, true},
		{"5 is accepted", 5, true},
		{"6 is not (divisible by 3)", 6, false},
		{"7 is accepted", 7, true},
		{"8 is accepted", 8, true},
		{"9 is not (divisible by 3)", 9, false},
		{"10 is accepted", 10, true},
		{"11 is accepted", 11, true},
		{"12 is not (divisible by 3)", 12, false},
		{"13 is not (ends with 3)", 13, false},
		{"14 is accepted", 14, true},
		{"23 is not (ends with 3)", 23, false},
		{"33 is not (divisible by 3 and ends with 3)", 33, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAccepted(tt.input)
			if result != tt.expected {
				t.Errorf("isAccepted(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
