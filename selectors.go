package filter

// and is a convenience function for combining two bools with AND
func and(a, b bool) bool {
	return a && b
}

// or is a convenience function for combining two bools with OR
func or(a, b bool) bool {
	return a || b
}

// selector is a convenience function for combining multiple filters with a combine function
func selector[T any](initialValue bool, combine func(bool, bool) bool, filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		appendLocation := initialValue
		for _, filter := range filters {
			appendLocation = combine(appendLocation, filter(val))
		}
		return appendLocation
	}
}

func And[T any](filters ...func(T) bool) func(T) bool {
	return selector(true, and, filters...)
}

// andSelector is a convenience function for combining multiple filters with AND
func andSelector[T any](filters ...func(T) bool) func(T) bool {
	return selector(true, and, filters...)
}

// orSelector is a convenience function for combining multiple filters with OR
func orSelector[T any](filters ...func(T) bool) func(T) bool {
	return selector(false, or, filters...)
}

// notAndSelector is a convenience function for combining multiple filters with NOT AND
func notAndSelector[T any](filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		return !selector(true, and, filters...)(val)
	}
}

func notOrSelector[T any](filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		return !selector(false, or, filters...)(val)
	}
}
