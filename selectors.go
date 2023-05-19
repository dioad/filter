package filter

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

func selector[T any](initialValue bool, combine func(bool, bool) bool, filters ...func(T) bool) func(T) bool {
	return func(val T) bool {
		appendLocation := initialValue
		for _, filter := range filters {
			appendLocation = combine(appendLocation, filter(val))
		}
		return appendLocation
	}
}

func andSelector[T any](filters ...func(T) bool) func(T) bool {
	return selector(true, and, filters...)
}

func orSelector[T any](filters ...func(T) bool) func(T) bool {
	return selector(false, or, filters...)
}

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
