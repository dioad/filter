package filter

import (
	"errors"
)

var (
	ErrNoElement       = errors.New("no element")
	ErrTooManyElements = errors.New("too many elements")
)

// Equals is a convenience function for comparing two values
func Equals[T comparable](v T) func(T) bool {
	return func(a T) bool {
		return v == a
	}
}

// FilterSliceAnd is a convenience function for filtering a slice with multiple filters and combining the results with AND
func FilterSliceAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, andSelector(filters...))
}

// FilterSliceOr is a convenience function for filtering a slice with multiple filters and combining the results with OR
func FilterSliceOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, orSelector(filters...))
}

// FilterSliceNotOr is a convenience function for filtering a slice with multiple filters and combining the results with NOT OR
func FilterSliceNotOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notOrSelector(filters...))
}

// FilterSliceNotAnd is a convenience function for filtering a slice with multiple filters and combining the results with NOT AND
func FilterSliceNotAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notAndSelector(filters...))
}

// FilterSlice is a convenience function for filtering a slice with a single filter
func FilterSlice[T any](l []T, filter func(T) bool) []T {
	filtered := make([]T, 0)

	for _, val := range l {
		if filter(val) {
			filtered = append(filtered, val)
		}
	}

	return filtered
}

// OneOnly is a convenience function for getting the only element in a slice
func OneOnly[T any](l []T) (*T, error) {
	if len(l) == 0 {
		return nil, ErrNoElement
	} else if len(l) > 1 {
		return nil, ErrTooManyElements
	}
	return &l[0], nil
}

// SliceToOrFilter is a convenience function for converting a slice to a filter function
func SliceToOrFilter[T any, U any](list []T, filterFunc func(T) func(U) bool) func(U) bool {
	filters := make([]func(U) bool, 0)

	for _, item := range list {
		filters = append(filters, filterFunc(item))
	}

	return orSelector(filters...)
}

// SliceToAndFilter is a convenience function for converting a slice to a filter function
func SliceToAndFilter[T any, U any](list []T, filterFunc func(T) func(U) bool) func(U) bool {
	filters := make([]func(U) bool, 0)

	for _, item := range list {
		filters = append(filters, filterFunc(item))
	}

	return andSelector(filters...)
}
