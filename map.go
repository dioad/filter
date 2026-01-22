// Copyright 2024 Dioad Consulting Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter

// FilterMapAnd filters a map using multiple filter functions combined with logical AND.
// A key-value pair is included in the result only if all filters return true for the value.
//
// Parameters:
//   - m: The input map to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new map containing only the key-value pairs where the value passed all filters
//
// Example:
//
//	users := map[string]User{
//	    "alice": {Age: 25, Role: "Admin"},
//	    "bob":   {Age: 30, Role: "User"},
//	    "carol": {Age: 22, Role: "Admin"},
//	}
//	isAdmin := func(u User) bool { return u.Role == "Admin" }
//	isOver25 := func(u User) bool { return u.Age > 25 }
//	result := filter.FilterMapAnd(users, isAdmin, isOver25)
//	// result contains only "alice"
func FilterMapAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, andSelector(filters...))
}

// FilterMapOr filters a map using multiple filter functions combined with logical OR.
// A key-value pair is included in the result if any of the filters return true for the value.
//
// Parameters:
//   - m: The input map to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new map containing the key-value pairs where the value passed at least one filter
//
// Example:
//
//	users := map[string]User{
//	    "alice": {Age: 25, Role: "Admin"},
//	    "bob":   {Age: 30, Role: "User"},
//	    "carol": {Age: 22, Role: "Admin"},
//	}
//	isAdmin := func(u User) bool { return u.Role == "Admin" }
//	isOver25 := func(u User) bool { return u.Age > 25 }
//	result := filter.FilterMapOr(users, isAdmin, isOver25)
//	// result contains "alice", "bob", and "carol"
func FilterMapOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, orSelector(filters...))
}

// FilterMapNotOr filters a map using multiple filter functions combined with logical NOT OR.
// A key-value pair is included in the result only if all filters return false for the value.
//
// Parameters:
//   - m: The input map to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new map containing only the key-value pairs where the value failed all filters
//
// Example:
//
//	users := map[string]User{
//	    "alice": {Age: 25, Role: "Admin"},
//	    "bob":   {Age: 30, Role: "User"},
//	    "carol": {Age: 22, Role: "Admin"},
//	}
//	isAdmin := func(u User) bool { return u.Role == "Admin" }
//	isOver25 := func(u User) bool { return u.Age > 25 }
//	result := filter.FilterMapNotOr(users, isAdmin, isOver25)
//	// result contains only "bob" (not admin and not over 25)
func FilterMapNotOr[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notOrSelector(filters...))
}

// FilterMapNotAnd filters a map using multiple filter functions combined with logical NOT AND.
// A key-value pair is included in the result if any of the filters return false for the value.
//
// Parameters:
//   - m: The input map to filter
//   - filters: A variadic list of filter functions
//
// Returns:
//   - A new map containing the key-value pairs where the value failed at least one filter
//
// Example:
//
//	users := map[string]User{
//	    "alice": {Age: 25, Role: "Admin"},
//	    "bob":   {Age: 30, Role: "User"},
//	    "carol": {Age: 22, Role: "Admin"},
//	}
//	isAdmin := func(u User) bool { return u.Role == "Admin" }
//	isOver25 := func(u User) bool { return u.Age > 25 }
//	result := filter.FilterMapNotAnd(users, isAdmin, isOver25)
//	// result contains "bob" and "carol" (either not admin or not over 25)
func FilterMapNotAnd[K comparable, V any](m map[K]V, filters ...func(V) bool) map[K]V {
	return FilterMap(m, notAndSelector(filters...))
}

// FilterMap filters a map using a single filter function.
// A key-value pair is included in the result only if the filter returns true for the value.
//
// Parameters:
//   - m: The input map to filter
//   - filter: A function that takes a value and returns a boolean
//
// Returns:
//   - A new map containing only the key-value pairs where the value passed the filter
//
// Example:
//
//	users := map[string]int{
//	    "alice": 25,
//	    "bob":   30,
//	    "carol": 22,
//	}
//	isOver25 := func(age int) bool { return age > 25 }
//	result := filter.FilterMap(users, isOver25)
//	// result = map[bob:30]
func FilterMap[K comparable, V any](m map[K]V, filter func(V) bool) map[K]V {
	// Estimate the capacity as half of the original map size
	estimatedCapacity := len(m) / 2
	if estimatedCapacity < 1 {
		estimatedCapacity = len(m)
	}

	filtered := make(map[K]V, estimatedCapacity)

	for key, val := range m {
		if filter(val) {
			filtered[key] = val
		}
	}

	return filtered
}
