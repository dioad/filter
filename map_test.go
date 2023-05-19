package filter

import (
	"reflect"
	"testing"
)

func TestFilterMap(t *testing.T) {
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
