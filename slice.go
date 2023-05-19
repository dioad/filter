package filter

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
