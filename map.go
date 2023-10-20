package filter

// FilterMapAnd is a convenience function for filtering a map with multiple filters and combining the results with AND
func FilterMapAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, andSelector(filters...))
}

// FilterMapOr is a convenience function for filtering a map with multiple filters and combining the results with OR
func FilterMapOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, orSelector(filters...))
}

// FilterMapNotOr is a convenience function for filtering a map with multiple filters and combining the results with NOT OR
func FilterMapNotOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notOrSelector(filters...))
}

// FilterMapNotAnd is a convenience function for filtering a map with multiple filters and combining the results with NOT AND
func FilterMapNotAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notAndSelector(filters...))
}

// FilterMap is a convenience function for filtering a map with a single filter
func FilterMap[K comparable, V any](m map[K]V, filter func(V) bool) map[K]V {
	filtered := make(map[K]V)

	for key, val := range m {
		if filter(val) {
			filtered[key] = val
		}
	}

	return filtered
}
