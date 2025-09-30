package main

import (
	"testing"
)

func TestJellyTransformation(t *testing.T) {
	tests := []struct {
		name  string
		bears string
		K     int
		want  int
	}{
		{
			name:  "Example 1: IPYIYP with K=3",
			bears: "IPYIYP",
			K:     3,
			want:  3,
		},
		{
			name:  "Example 2: IY with K=1",
			bears: "IY",
			K:     1,
			want:  2,
		},
		{
			name:  "Example 3: PPY with K=2",
			bears: "PPY",
			K:     2,
			want:  -1,
		},
		{
			name:  "Already all ice",
			bears: "III",
			K:     2,
			want:  0,
		},
		{
			name:  "Single jelly transformation",
			bears: "Y",
			K:     1,
			want:  2,
		},
		{
			name:  "YY with K=2 - should be possible",
			bears: "YY",
			K:     2,
			want:  2, // YY -> PP -> II
		},
		{
			name:  "Empty string",
			bears: "",
			K:     1,
			want:  0,
		},
		{
			name:  "Single ice jelly",
			bears: "I",
			K:     1,
			want:  0,
		},
		{
			name:  "YPY with K=1 - correct calculation",
			bears: "YPY",
			K:     1,
			want:  5, // Y(2) + P(1) + Y(2) = 5
		},
		{
			name:  "YYYY with K=2 - should be possible",
			bears: "YYYY",
			K:     2,
			want:  4, // YYYY -> PPPP -> IIII
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jellyTransformation(tt.bears, tt.K)
			if got != tt.want {
				t.Errorf("jellyTransformation(%s, %d) = %d, want %d", tt.bears, tt.K, got, tt.want)
			}
		})
	}
}

func TestTransformSegment(t *testing.T) {
	tests := []struct {
		name     string
		bears    string
		start    int
		K        int
		expected string
	}{
		{
			name:     "Transform Y to P",
			bears:    "Y",
			start:    0,
			K:        1,
			expected: "P",
		},
		{
			name:     "Transform P to I",
			bears:    "P",
			start:    0,
			K:        1,
			expected: "I",
		},
		{
			name:     "Transform I to Y",
			bears:    "I",
			start:    0,
			K:        1,
			expected: "Y",
		},
		{
			name:     "Transform segment of 3",
			bears:    "YPY",
			start:    0,
			K:        3,
			expected: "PIP",
		},
		{
			name:     "Transform middle segment",
			bears:    "YYY",
			start:    1,
			K:        1,
			expected: "YPY",
		},
		{
			name:     "Transform end segment",
			bears:    "YYY",
			start:    2,
			K:        1,
			expected: "YYP",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformSegment(tt.bears, tt.start, tt.K)
			if got != tt.expected {
				t.Errorf("transformSegment(%s, %d, %d) = %s, want %s",
					tt.bears, tt.start, tt.K, got, tt.expected)
			}
		})
	}
}

func TestAllIce(t *testing.T) {
	tests := []struct {
		name     string
		bears    string
		expected bool
	}{
		{"All ice", "III", true},
		{"Mixed", "IPY", false},
		{"No ice", "YPY", false},
		{"Empty", "", true},
		{"Single ice", "I", true},
		{"Single yellow", "Y", false},
		{"Single pink", "P", false},
		{"Mixed with ice", "IYI", false},
		{"All yellow", "YYY", false},
		{"All pink", "PPP", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allIce(tt.bears)
			if got != tt.expected {
				t.Errorf("allIce(%s) = %t, want %t", tt.bears, got, tt.expected)
			}
		})
	}
}
