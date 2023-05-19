package filter

func FilterMapAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, andSelector(filters...))
}

func FilterMapOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, orSelector(filters...))
}

func FilterMapNotOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notOrSelector(filters...))
}

func FilterMapNotAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notAndSelector(filters...))
}

func FilterMap[K comparable, V any](m map[K]V, filter func(V) bool) map[K]V {
	filtered := make(map[K]V)

	for key, val := range m {
		if filter(val) {
			filtered[key] = val
		}
	}

	return filtered
}
