package utils

import (
	"testing"
)

func TestRomans_ToInt(t *testing.T) {
	cases := []struct{
		testName string
		romanNumeral string
		expected int
	} {
		{
			testName: "1. Positive Case: Valid roman, less than 50",
			romanNumeral: "XXIV",
			expected: 24,
		},
		{
			testName: "2. Positive Case: Invalid roman or more than 50",
			romanNumeral: "CW",
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		romans := NewRomans()
		result := romans.ToInt(c.romanNumeral)
		if c.expected != result {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, result)
		}
	}
}
