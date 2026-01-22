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

package filter_test

import (
	"fmt"

	"github.com/dioad/filter"
)

// This file contains examples of how to use the filter package.
// These examples serve as both documentation and tests.

func Example_filteringSlices() {
	// Basic filtering
	numbers := []int{1, 2, 3, 4, 5}

	// Filter even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	evenNumbers := filter.FilterSlice(numbers, isEven)
	fmt.Println("Even numbers:", evenNumbers)

	// Combining filters with AND
	isGreaterThanThree := func(n int) bool {
		return n > 3
	}

	evenAndGreaterThanThree := filter.FilterSliceAnd(numbers, isEven, isGreaterThanThree)
	fmt.Println("Even and greater than 3:", evenAndGreaterThanThree)

	// Combining filters with OR
	isLessThanTwo := func(n int) bool {
		return n < 2
	}

	evenOrLessThanTwo := filter.FilterSliceOr(numbers, isEven, isLessThanTwo)
	fmt.Println("Even or less than 2:", evenOrLessThanTwo)

	// Output:
	// Even numbers: [2 4]
	// Even and greater than 3: [4]
	// Even or less than 2: [1 2 4]
}

func Example_filteringMaps() {
	// Map filtering
	users := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
		"Dave":  35,
	}

	// Filter users older than 25
	isOlderThan25 := func(age int) bool {
		return age > 25
	}

	olderUsers := filter.FilterMap(users, isOlderThan25)
	fmt.Println("Users older than 25:", olderUsers)

	// Filter users younger than 30 and older than 20
	isYoungerThan30 := func(age int) bool {
		return age < 30
	}
	isOlderThan20 := func(age int) bool {
		return age > 20
	}

	middleAgedUsers := filter.FilterMapAnd(users, isYoungerThan30, isOlderThan20)
	fmt.Println("Users between 20 and 30:", middleAgedUsers)

	// Output:
	// Users older than 25: map[Bob:30 Dave:35]
	// Users between 20 and 30: map[Alice:25 Carol:22]
}

func Example_selectingFromSlices() {
	fruits := []string{"apple", "banana", "cherry", "date"}

	// Select first element
	first := filter.SliceSelectFirst(fruits)
	fmt.Println("First fruit:", *first)

	// Select last element
	last := filter.SliceSelectLast(fruits)
	fmt.Println("Last fruit:", *last)

	// Select a specific element using filtering
	cherries := filter.FilterSlice(fruits, filter.Equals("cherry"))
	cherry, err := filter.OneOnly(cherries)
	if err == nil {
		fmt.Println("Found cherry:", *cherry)
	}

	// Output:
	// First fruit: apple
	// Last fruit: date
	// Found cherry: cherry
}

func Example_combiningFilters() {
	// Creating reusable filters
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	isGreaterThan10 := func(n int) bool { return n > 10 }

	// Combining filters with logical operations
	isEvenAndPositive := filter.And(isEven, isPositive)
	isEvenOrGreaterThan10 := filter.Or(isEven, isGreaterThan10)
	isNotEvenAndNotPositive := filter.NotAnd(isEven, isPositive)

	numbers := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 11, 12}

	// Apply the combined filters
	evenAndPositive := filter.FilterSlice(numbers, isEvenAndPositive)
	evenOrGreaterThan10 := filter.FilterSlice(numbers, isEvenOrGreaterThan10)
	notEvenAndNotPositive := filter.FilterSlice(numbers, isNotEvenAndNotPositive)

	fmt.Println("Even and positive:", evenAndPositive)
	fmt.Println("Even or greater than 10:", evenOrGreaterThan10)
	fmt.Println("Not (even and positive):", notEvenAndNotPositive)

	// Output:
	// Even and positive: [2 4 12]
	// Even or greater than 10: [-4 -2 0 2 4 11 12]
	// Not (even and positive): [-4 -3 -2 -1 0 1 3 11]
}

func Example_convertingSlicesToFilters() {
	// Define a list of allowed values
	allowedIDs := []int{1, 3, 5}

	// Convert the slice to a filter function
	hasAllowedID := filter.SliceToOrFilter(allowedIDs, func(id int) func(int) bool {
		return filter.Equals(id)
	})

	// Apply the filter to a list of IDs
	allIDs := []int{1, 2, 3, 4, 5, 6}
	filteredIDs := filter.FilterSlice(allIDs, hasAllowedID)

	fmt.Println("Allowed IDs:", filteredIDs)

	// Output:
	// Allowed IDs: [1 3 5]
}
