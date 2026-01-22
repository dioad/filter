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

// Package filter provides generic filtering utilities for maps and slices with various logical operations.
//
// The package is designed to make filtering operations on collections more concise and readable
// by providing a set of utility functions that can be combined to create complex filtering logic.
//
// Key features:
//
// - Filter maps and slices with simple predicates
// - Combine multiple filters with logical operations (AND, OR, NOT AND, NOT OR)
// - Select elements from slices (first, last, random)
// - Convert slices to filter functions
// - Type-safe operations using Go generics
//
// Example usage for filtering slices:
//
//	numbers := []int{1, 2, 3, 4, 5}
//
//	// Filter even numbers
//	isEven := func(n int) bool {
//	    return n%2 == 0
//	}
//
//	evenNumbers := filter.FilterSlice(numbers, isEven)
//	// evenNumbers = [2, 4]
//
// Example usage for filtering maps:
//
//	users := map[string]int{
//	    "Alice": 25,
//	    "Bob":   30,
//	    "Carol": 22,
//	    "Dave":  35,
//	}
//
//	// Filter users older than 25
//	isOlderThan25 := func(age int) bool {
//	    return age > 25
//	}
//
//	olderUsers := filter.FilterMap(users, isOlderThan25)
//	// olderUsers = map[Bob:30 Dave:35]
//
// For more examples and detailed documentation, see the README.md file.
package filter
