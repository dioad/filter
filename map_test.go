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

import (
	"reflect"
	"testing"
)

func TestFilterMap(t *testing.T) {
	t.Run("basic filtering", func(t *testing.T) {
		input := map[int]string{
			1: "a",
			2: "b",
			3: "c",
		}
		expected := map[int]string{
			3: "c",
		}

		filtered := FilterMap(input, match("c"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int]string{}
		expected := map[int]string{}

		filtered := FilterMap(input, match("c"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("no matches", func(t *testing.T) {
		input := map[int]string{
			1: "a",
			2: "b",
			3: "c",
		}
		expected := map[int]string{}

		filtered := FilterMap(input, match("d"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("all matches", func(t *testing.T) {
		input := map[int]string{
			1: "a",
			2: "a",
			3: "a",
		}
		expected := map[int]string{
			1: "a",
			2: "a",
			3: "a",
		}

		filtered := FilterMap(input, match("a"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})
}

func TestFilterMapAnd(t *testing.T) {
	input := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	expected := map[string]string{
		"c": "c",
	}

	filtered := FilterMapAnd(input, match("c"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestFilterMapOr(t *testing.T) {
	input := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	expected := map[string]string{
		"b": "b",
		"c": "c",
	}

	filtered := FilterMapOr(input, match("b"), match("c"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestFilterMapNotOr(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	expected := map[string]int{
		"a": 1,
	}

	filtered := FilterMapNotOr(input, match(2), match(3))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestFilterMapNotAnd(t *testing.T) {
	input := map[string]complexType{
		"a": complexType{A: 1, B: "a"},
		"b": complexType{A: 2, B: "b"},
		"c": complexType{A: 3, B: "c"},
	}
	expected := map[string]complexType{
		"b": complexType{A: 2, B: "b"},
		"c": complexType{A: 3, B: "c"},
	}

	filtered := FilterMapNotAnd(input, matchA(1), matchB("a"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}
