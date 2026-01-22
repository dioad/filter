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
	"errors"
)

var (
	// ErrNoElement is returned when an operation expects at least one element but the slice is empty
	ErrNoElement = errors.New("no element")

	// ErrTooManyElements is returned when an operation expects exactly one element but the slice has more
	ErrTooManyElements = errors.New("too many elements")
)

// Equals creates a filter function that checks if a value equals the provided value.
// This is a basic building block for creating more complex filters.
//
// Example:
//
//	isThree := filter.Equals(3)
//	// isThree(3) returns true
//	// isThree(4) returns false
func Equals[T comparable](v T) func(T) bool {
	return func(a T) bool {
		return v == a
	}
}

// FilterSliceAnd filters a slice using multiple filter functions combined with logical AND.
// An element is included in the result only if all filters return true for it.
//
// Parameters:
//   - l: The input slice to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new slice containing only the elements that passed all filters
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	isGreaterThanThree := func(n int) bool { return n > 3 }
//	result := filter.FilterSliceAnd(numbers, isEven, isGreaterThanThree)
//	// result = [4]
func FilterSliceAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, andSelector(filters...))
}

// FilterSliceOr filters a slice using multiple filter functions combined with logical OR.
// An element is included in the result if any of the filters return true for it.
//
// Parameters:
//   - l: The input slice to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new slice containing the elements that passed at least one filter
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	isLessThanTwo := func(n int) bool { return n < 2 }
//	result := filter.FilterSliceOr(numbers, isEven, isLessThanTwo)
//	// result = [1, 2, 4]
func FilterSliceOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, orSelector(filters...))
}

// FilterSliceNotOr filters a slice using multiple filter functions combined with logical NOT OR.
// An element is included in the result only if all filters return false for it.
//
// Parameters:
//   - l: The input slice to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new slice containing only the elements for which all filters returned false
func FilterSliceNotOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notOrSelector(filters...))
}

// FilterSliceNotAnd filters a slice using multiple filter functions combined with logical NOT AND.
// An element is included in the result if any of the filters return false for it.
//
// Parameters:
//   - l: The input slice to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new slice containing the elements for which at least one filter returned false
func FilterSliceNotAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notAndSelector(filters...))
}

// FilterSlice filters a slice using a single filter function.
// An element is included in the result only if the filter returns true for it.
//
// Parameters:
//   - l: The input slice to filter
//   - filter: A function that takes an element and returns a boolean
//
// Returns:
//   - A new slice containing only the elements that passed the filter
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	result := filter.FilterSlice(numbers, isEven)
//	// result = [2, 4]
func FilterSlice[T any](l []T, filter func(T) bool) []T {
	filtered := make([]T, 0, len(l)/2) // Pre-allocate with a reasonable capacity

	for _, val := range l {
		if filter(val) {
			filtered = append(filtered, val)
		}
	}

	return filtered
}

// OneOnly returns the only element in a slice, or an error if the slice doesn't contain exactly one element.
//
// Parameters:
//   - l: The input slice
//
// Returns:
//   - A pointer to the only element in the slice, or nil if an error occurred
//   - An error: ErrNoElement if the slice is empty, ErrTooManyElements if it has more than one element
//
// Example:
//
//	numbers := []int{42}
//	result, err := filter.OneOnly(numbers)
//	// result points to 42, err is nil
func OneOnly[T any](l []T) (*T, error) {
	if len(l) == 0 {
		return nil, ErrNoElement
	} else if len(l) > 1 {
		return nil, ErrTooManyElements
	}
	return &l[0], nil
}

// SliceToOrFilter converts a slice of values into a single filter function using logical OR.
// The returned filter function returns true if any of the individual filters (created from the slice elements) return true.
//
// Parameters:
//   - list: A slice of values to convert into filters
//   - filterFunc: A function that converts each value in the list to a filter function
//
// Returns:
//   - A filter function that combines all the individual filters with logical OR
//
// Example:
//
//	allowedIDs := []int{1, 3, 5}
//	hasAllowedID := filter.SliceToOrFilter(allowedIDs, func(id int) func(user User) bool {
//	    return func(user User) bool { return user.ID == id }
//	})
//	// hasAllowedID returns true for users with ID 1, 3, or 5
func SliceToOrFilter[T any, U any](list []T, filterFunc func(T) func(U) bool) func(U) bool {
	filters := make([]func(U) bool, 0, len(list))

	for _, item := range list {
		filters = append(filters, filterFunc(item))
	}

	return orSelector(filters...)
}

// SliceToAndFilter converts a slice of values into a single filter function using logical AND.
// The returned filter function returns true only if all of the individual filters (created from the slice elements) return true.
//
// Parameters:
//   - list: A slice of values to convert into filters
//   - filterFunc: A function that converts each value in the list to a filter function
//
// Returns:
//   - A filter function that combines all the individual filters with logical AND
//
// Example:
//
//	requiredTags := []string{"important", "urgent"}
//	hasAllRequiredTags := filter.SliceToAndFilter(requiredTags, func(tag string) func(task Task) bool {
//	    return func(task Task) bool { return task.HasTag(tag) }
//	})
//	// hasAllRequiredTags returns true only for tasks that have both "important" and "urgent" tags
func SliceToAndFilter[T any, U any](list []T, filterFunc func(T) func(U) bool) func(U) bool {
	filters := make([]func(U) bool, 0, len(list))

	for _, item := range list {
		filters = append(filters, filterFunc(item))
	}

	return andSelector(filters...)
}
