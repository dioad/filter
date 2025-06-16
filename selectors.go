package filter

// and is a helper function for combining two boolean values with logical AND operation.
// It returns true only if both a and b are true.
func and(a, b bool) bool {
	return a && b
}

// or is a helper function for combining two boolean values with logical OR operation.
// It returns true if either a or b (or both) are true.
func or(a, b bool) bool {
	return a || b
}

// selector is a higher-order function that creates a filter by combining multiple filter functions.
// It takes an initial value, a combine function (like and/or), and a variadic list of filter functions.
// The returned function applies all filters to a value and combines the results using the combine function.
//
// Parameters:
//   - initialValue: The starting value for the combination (true for AND, false for OR)
//   - combine: A function that combines two boolean values (like and/or)
//   - filters: A variadic list of filter functions to be combined
//
// Returns:
//   - A function that takes a value of type T and returns a boolean result
func selector[T any](initialValue bool, combine func(bool, bool) bool, filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		result := initialValue
		for _, filter := range filters {
			result = combine(result, filter(val))
		}
		return result
	}
}

// And creates a filter function that combines multiple filters with logical AND.
// The returned filter returns true only if all the provided filters return true for a given value.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	isPositive := func(n int) bool { return n > 0 }
//	isEvenAndPositive := filter.And(isEven, isPositive)
//	// isEvenAndPositive(4) returns true
//	// isEvenAndPositive(-2) returns false
func And[T any](filters ...func(T) bool) func(T) bool {
	return selector(true, and, filters...)
}

// andSelector is an alias for And, used internally for consistent naming with other selector functions.
// This is an internal function used by other package functions.
func andSelector[T any](filters ...func(T) bool) func(T) bool {
	return And(filters...)
}

// Or creates a filter function that combines multiple filters with logical OR.
// The returned filter returns true if any of the provided filters return true for a given value.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	isNegative := func(n int) bool { return n < 0 }
//	isEvenOrNegative := filter.Or(isEven, isNegative)
//	// isEvenOrNegative(4) returns true
//	// isEvenOrNegative(-3) returns true
//	// isEvenOrNegative(3) returns false
func Or[T any](filters ...func(T) bool) func(T) bool {
	return selector(false, or, filters...)
}

// NotAnd creates a filter function that combines multiple filters with logical NOT AND.
// The returned filter returns true only if at least one of the provided filters returns false for a given value.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	isPositive := func(n int) bool { return n > 0 }
//	isNotEvenOrNotPositive := filter.NotAnd(isEven, isPositive)
//	// isNotEvenOrNotPositive(4) returns false
//	// isNotEvenOrNotPositive(-2) returns true
//	// isNotEvenOrNotPositive(3) returns true
func NotAnd[T any](filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		return !And(filters...)(val)
	}
}

// NotOr creates a filter function that combines multiple filters with logical NOT OR.
// The returned filter returns true only if all of the provided filters return false for a given value.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	isNegative := func(n int) bool { return n < 0 }
//	isNotEvenAndNotNegative := filter.NotOr(isEven, isNegative)
//	// isNotEvenAndNotNegative(4) returns false
//	// isNotEvenAndNotNegative(-3) returns false
//	// isNotEvenAndNotNegative(3) returns true
func NotOr[T any](filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		return !Or(filters...)(val)
	}
}

// orSelector is an alias for Or, used internally for consistent naming with other selector functions.
// This is an internal function used by other package functions.
func orSelector[T any](filters ...func(T) bool) func(T) bool {
	return Or(filters...)
}

// notAndSelector is an alias for NotAnd, used internally for consistent naming with other selector functions.
// This is an internal function used by other package functions.
func notAndSelector[T any](filters ...func(T) bool) func(T) bool {
	return NotAnd(filters...)
}

// notOrSelector is an alias for NotOr, used internally for consistent naming with other selector functions.
// This is an internal function used by other package functions.
func notOrSelector[T any](filters ...func(T) bool) func(T) bool {
	return NotOr(filters...)
}
