package filter

import (
	"testing"
)

func TestAnd(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	isEvenAndPositive := And(isEven, isPositive)

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even and positive", 4, true},
		{"odd and positive", 3, false},
		{"even and negative", -2, false},
		{"odd and negative", -3, false},
		{"zero", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isEvenAndPositive(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}
}

func TestOr(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isNegative := func(n int) bool { return n < 0 }
	isEvenOrNegative := Or(isEven, isNegative)

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even and positive", 4, true},
		{"odd and positive", 3, false},
		{"even and negative", -2, true},
		{"odd and negative", -3, true},
		{"zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isEvenOrNegative(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}
}

func TestNotAnd(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	isNotEvenOrNotPositive := NotAnd(isEven, isPositive)

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even and positive", 4, false},
		{"odd and positive", 3, true},
		{"even and negative", -2, true},
		{"odd and negative", -3, true},
		{"zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isNotEvenOrNotPositive(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}
}

func TestNotOr(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isNegative := func(n int) bool { return n < 0 }
	isNotEvenAndNotNegative := NotOr(isEven, isNegative)

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even and positive", 4, false},
		{"odd and positive", 3, true},
		{"even and negative", -2, false},
		{"odd and negative", -3, false},
		{"zero", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isNotEvenAndNotNegative(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}
}

func TestEmptyFilters(t *testing.T) {
	t.Run("And with no filters", func(t *testing.T) {
		alwaysTrue := And[int]()
		if !alwaysTrue(42) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("Or with no filters", func(t *testing.T) {
		alwaysFalse := Or[int]()
		if alwaysFalse(42) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("NotAnd with no filters", func(t *testing.T) {
		alwaysFalse := NotAnd[int]()
		if alwaysFalse(42) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("NotOr with no filters", func(t *testing.T) {
		alwaysTrue := NotOr[int]()
		if !alwaysTrue(42) {
			t.Errorf("expected true, got false")
		}
	})
}
