package models

import (
	"sort"
	"testing"
)

type MockRoyalNames struct {
	MockLen  func() int
	MockLess func(i, j int) bool
	MockSwap func(i, j int)
}

func (mrn MockRoyalNames) Len() int {
	return mrn.MockLen()
}

func (mrn MockRoyalNames) Less(i, j int) bool {
	return mrn.MockLess(i, j)
}

func (mrn MockRoyalNames) Swap(i, j int) {
	mrn.MockSwap(i, j)
}

func TestRoyalNames_Len(t *testing.T) {
	cases := []struct {
		testName   string
		royalNames RoyalNames
		expected   int
	}{
		{
			testName:   "1. Positive Case",
			royalNames: RoyalNames{"David IV", "William II"},
			expected:   2,
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		result := c.royalNames.Len()
		if c.expected != result {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, result)
		}
	}
}

func TestRoyalNames_Less(t *testing.T) {
	cases := []struct {
		testName   string
		royalNames RoyalNames
		expected   bool
	}{
		{
			testName:   "1. Positive Case: David IV is less than William II",
			royalNames: RoyalNames{"David IV", "William II"},
			expected:   true,
		},
		{
			testName:   "2. Positive Case: William II is not less than David IV",
			royalNames: RoyalNames{"William II", "David IV"},
			expected:   false,
		},
		{
			testName:   "3. Positive Case: William II is not less than William I",
			royalNames: RoyalNames{"William II", "William I"},
			expected:   false,
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		result := c.royalNames.Less(0, 1)
		if c.expected != result {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, result)
		}
	}
}

func TestRoyalNames_Swap(t *testing.T) {
	cases := []struct {
		testName   string
		royalNames RoyalNames
		expected   RoyalNames
	}{
		{
			testName:   "1. Positive Case",
			royalNames: RoyalNames{"David IV", "William II"},
			expected:   RoyalNames{"William II", "David IV"},
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		c.royalNames.Swap(0, 1)
		if c.expected[0] != c.royalNames[0] {
			t.Errorf("Expected: %v, Got: %v\n", c.expected[0], c.royalNames[1])
		}
	}
}

func TestRoyalNames_Sort(t *testing.T) {
	cases := []struct {
		testName   string
		royalNames RoyalNames
		expected   RoyalNames
	}{
		{
			testName:   "1. Positive Case: Different first name",
			royalNames: RoyalNames{"William II", "David IV"},
			expected:   RoyalNames{"David IV", "William II"},
		},
		{
			testName:   "2. Positive Case: Same first name",
			royalNames: RoyalNames{"William IV", "William II"},
			expected:   RoyalNames{"William II", "William IV"},
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		sort.Sort(c.royalNames)
		if c.expected[0] != c.royalNames[0] {
			t.Errorf("Expected: %v, Got: %v\n", c.expected[0], c.royalNames[1])
		}
	}
}
