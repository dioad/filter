package filter

import (
	"errors"
)

var (
	ErrNoElement       = errors.New("no element")
	ErrTooManyElements = errors.New("too many elements")
)

func Equals[T comparable](v T) func(T) bool {
	return func(a T) bool {
		return v == a
	}
}

func FilterSliceAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, andSelector(filters...))
}

func FilterSliceOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, orSelector(filters...))
}

func FilterSliceNotOr[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notOrSelector(filters...))
}

func FilterSliceNotAnd[T any](l []T, filters ...func(T) bool) []T {
	return FilterSlice(l, notAndSelector(filters...))
}

func FilterSlice[T any](l []T, filter func(T) bool) []T {
	filtered := make([]T, 0)

	for _, val := range l {
		if filter(val) {
			filtered = append(filtered, val)
		}
	}

	return filtered
}

func OneOnly[T any](l []T) (*T, error) {
	if len(l) == 0 {
		return nil, ErrNoElement
	} else if len(l) > 1 {
		return nil, ErrTooManyElements
	} else {
		return &l[0], nil
	}
}
