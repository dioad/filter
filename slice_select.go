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
	"math/rand"
)

// SliceSelectFirst returns a pointer to the first element in a slice.
// This is useful when you need to access the first element without modifying the original slice.
//
// Parameters:
//   - l: The input slice
//
// Returns:
//   - A pointer to the first element, or nil if the slice is empty
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	first := filter.SliceSelectFirst(items)
//	// *first = "apple"
func SliceSelectFirst[T any](l []T) *T {
	if len(l) == 0 {
		return nil
	}

	return &l[0]
}

// SliceSelectLast returns a pointer to the last element in a slice.
// This is useful when you need to access the last element without modifying the original slice.
//
// Parameters:
//   - l: The input slice
//
// Returns:
//   - A pointer to the last element, or nil if the slice is empty
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	last := filter.SliceSelectLast(items)
//	// *last = "cherry"
func SliceSelectLast[T any](l []T) *T {
	if len(l) == 0 {
		return nil
	}

	return &l[len(l)-1]
}

// SliceSelectRandom returns a pointer to a random element in a slice.
// This function uses the default random number generator.
//
// Note: This function is not cryptographically secure and should not be used
// for security-sensitive applications.
//
// Parameters:
//   - l: The input slice
//
// Returns:
//   - A pointer to a randomly selected element, or nil if the slice is empty
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	random := filter.SliceSelectRandom(items)
//	// *random = one of "apple", "banana", or "cherry"
func SliceSelectRandom[T any](l []T) *T {
	return SliceSelectRandomWithGenerator(l, rand.Intn)
}

// SliceSelectRandomWithGenerator returns a pointer to a random element in a slice using a custom random number generator.
// This allows for deterministic random selection, which can be useful for testing.
//
// Parameters:
//   - l: The input slice
//   - generator: A function that takes the length of the slice and returns a random index
//
// Returns:
//   - A pointer to a randomly selected element, or nil if the slice is empty or if generator is nil
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	// Always select the first element for testing
//	alwaysFirst := func(n int) int { return 0 }
//	result := filter.SliceSelectRandomWithGenerator(items, alwaysFirst)
//	// *result = "apple"
func SliceSelectRandomWithGenerator[T any](l []T, generator func(int) int) *T {
	if len(l) == 0 {
		return nil
	}

	if generator != nil {
		return &l[generator(len(l))]
	}

	return nil
}
