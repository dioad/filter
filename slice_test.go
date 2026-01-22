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
	"errors"
	"reflect"
	"testing"
)

func match[T any](v T) func(T) bool {
	return func(a T) bool {
		return reflect.DeepEqual(a, v)
	}
}

func TestMatch(t *testing.T) {
	if !match("a")("a") {
		t.Errorf("expected true, got false")
	}
}

func TestFilterSliceEmptyFilter(t *testing.T) {
	t.Run("with AND operation", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{"a", "b", "c"}

		filtered := FilterSliceAnd(input, []func(string) bool{}...)

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("with OR operation", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{}

		filtered := FilterSliceOr(input, []func(string) bool{}...)

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})
}

func TestFilterSlice(t *testing.T) {
	t.Run("basic filtering", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{"c"}

		filtered := FilterSlice(input, match("c"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		expected := []string{}

		filtered := FilterSlice(input, match("c"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("no matches", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{}

		filtered := FilterSlice(input, match("d"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})

	t.Run("all matches", func(t *testing.T) {
		input := []string{"a", "a", "a"}
		expected := []string{"a", "a", "a"}

		filtered := FilterSlice(input, match("a"))

		if !reflect.DeepEqual(expected, filtered) {
			t.Errorf("expected %v, got %v", expected, filtered)
		}
	})
}

func TestFilterSliceAnd(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := []string{"c"}

	filtered := FilterSliceAnd(input, match("c"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestFilterSliceOr(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := []string{"b", "c"}

	filtered := FilterSliceOr(input, match("b"), match("c"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestFilterSliceNotOr(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1}

	filtered := FilterSliceNotOr(input, match(2), match(3))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

type complexType struct {
	A int
	B string
}

func matchA(a int) func(complexType) bool {
	return func(val complexType) bool {
		return val.A == a
	}
}

func matchB(b string) func(complexType) bool {
	return func(val complexType) bool {
		return val.B == b
	}
}

func TestFilterSliceNotAnd(t *testing.T) {
	input := []complexType{
		{A: 1, B: "a"},
		{A: 2, B: "b"},
		{A: 3, B: "c"},
	}
	expected := []complexType{
		{A: 2, B: "b"},
		{A: 3, B: "c"},
	}

	filtered := FilterSliceNotAnd(input, matchA(1), matchB("a"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func aFunc(s string) func(string) bool {
	return func(val string) bool {
		return val == s
	}
}

func TestSliceToOrFilterSimple(t *testing.T) {
	input := []string{"a", "b", "c"}
	filterValues := []string{"b", "c"}
	expected := []string{"b", "c"}

	filter := SliceToOrFilter(filterValues, aFunc)

	filtered := FilterSlice(input, filter)

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestSliceToOrFilter(t *testing.T) {
	input := []sliceToFilterTestStruct{
		{input: []string{"a", "b", "c"}},
		{input: []string{"a", "c", "d"}},
		{input: []string{"a", "c"}},
	}
	filterValues := []string{"b", "d"}
	expected := []sliceToFilterTestStruct{
		{input: []string{"a", "b", "c"}},
		{input: []string{"a", "c", "d"}},
	}

	filter := SliceToOrFilter(filterValues, sliceToFilterTestFunc)

	filtered := FilterSlice(input, filter)

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

type sliceToFilterTestStruct struct {
	input []string
}

func sliceToFilterTestFunc(test string) func(sliceToFilterTestStruct) bool {
	return func(val sliceToFilterTestStruct) bool {
		filtered := FilterSlice(val.input, Equals(test))

		return len(filtered) > 0
	}
}

func TestSliceToAndFilter(t *testing.T) {
	input := []sliceToFilterTestStruct{
		{input: []string{"a", "b", "c"}},
		{input: []string{"a", "c", "d"}},
	}
	filterValues := []string{"b", "c"}
	expected := []sliceToFilterTestStruct{
		{input: []string{"a", "b", "c"}},
	}

	filter := SliceToAndFilter(filterValues, sliceToFilterTestFunc)

	filtered := FilterSlice(input, filter)

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestOneOnly(t *testing.T) {
	input := []string{"a"}
	expected := "a"

	filtered, err := OneOnly(input)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if *filtered != expected {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
}

func TestOneOnlyNoElements(t *testing.T) {
	input := []string{}
	expected := ErrNoElement

	_, err := OneOnly(input)

	if !errors.Is(err, expected) {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestOneOnlyTooManyElements(t *testing.T) {
	input := []string{"a", "b"}
	expected := ErrTooManyElements

	_, err := OneOnly(input)

	if !errors.Is(err, expected) {
		t.Errorf("expected %v, got %v", expected, err)
	}
}
