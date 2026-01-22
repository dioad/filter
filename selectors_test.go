// Copyright 2024 Dioad Consulting Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
