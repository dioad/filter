package filter

import (
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

func TestFilterSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := []string{"c"}

	filtered := FilterSlice(input, match("c"))

	if !reflect.DeepEqual(expected, filtered) {
		t.Errorf("expected %v, got %v", expected, filtered)
	}
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