package filter

// TODO: decide if the following should live in a separate module?

import (
	"math/rand"
	"time"
)

var (
	// DefaultRNG is the default random number generator
	DefaultRNG = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// SliceSelectFirst is a convenience function for selecting the first element from a slice
//
// returns nil if slice is empty
func SliceSelectFirst[T any](l []T) *T {
	if len(l) == 0 {
		return nil
	}

	return &l[0]
}

// SliceSelectLast is a convenience function for selecting the last element from a slice
//
// returns nil if slice is empty
func SliceSelectLast[T any](l []T) *T {
	if len(l) == 0 {
		return nil
	}

	return &l[len(l)-1]
}

// SliceSelectRandom is a convenience function for selecting a random element from a slice
//
// returns nil if slice is empty
func SliceSelectRandom[T any](l []T) *T {
	return SliceSelectRandomWithGenerator(l, DefaultRNG.Intn)
}

// SliceSelectRandomWithGenerator is a convenience function for selecting a random element from a slice with a custom random number generator
// The generator expects the length of the slice and should return a random number between 0 and length-1
//
// returns nil if slice is empty or if generator is nil
func SliceSelectRandomWithGenerator[T any](l []T, generator func(int) int) *T {
	if len(l) == 0 {
		return nil
	}

	if generator != nil {
		return &l[generator(len(l))]
	}

	return nil
}
