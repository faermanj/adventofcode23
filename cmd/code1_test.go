package cmd

import "testing"

func TestToNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{{"123abc456", 16},
		{"abc123", 13},
		{"a3abc", 33},
		{"", 0},
		{"abc", 0},
	}

	for _, tc := range testCases {
		result := toNumber(tc.input)
		if result != tc.expected {
			t.Errorf("toNumber(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
