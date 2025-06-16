package filter

import (
	"math/rand"
	"time"
)

var (
	// DefaultRNG is the default random number generator used by SliceSelectRandom.
	// It is initialized with a time-based seed for better randomness.
	DefaultRNG = rand.New(rand.NewSource(time.Now().UnixNano()))
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
// This function uses the DefaultRNG random number generator.
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
	return SliceSelectRandomWithGenerator(l, DefaultRNG.Intn)
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
