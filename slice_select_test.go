package filter

import (
	"slices"
	"testing"
)

func TestSliceSelectFirst(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "a"

	selected := SliceSelectFirst(input)

	if *selected != expected {
		t.Errorf("expected %v, got %v", expected, selected)
	}
}

func TestSliceSelectFirstWithEmptySlice(t *testing.T) {
	selected := SliceSelectFirst([]string{})

	if selected != nil {
		t.Errorf("expected nil, got %v", selected)
	}
}

func TestSliceSelectLast(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "c"

	selected := SliceSelectLast(input)

	if *selected != expected {
		t.Errorf("expected %v, got %v", expected, selected)
	}
}

func TestSliceSelectLastWithEmptySlice(t *testing.T) {
	selected := SliceSelectLast([]string{})

	if selected != nil {
		t.Errorf("expected nil, got %v", selected)
	}
}

func TestSliceSelectRandom(t *testing.T) {
	input := []string{"a", "b", "c"}

	selected := SliceSelectRandom(input)

	if !slices.Contains(input, *selected) {
		t.Errorf("expected %v to contain %v", input, *selected)
	}
}

func TestSliceSelectRandomWithGenerator(t *testing.T) {
	input := []string{"a", "b", "c"}

	selected := SliceSelectRandomWithGenerator(input, func(int) int { return 0 })

	if !slices.Contains(input, *selected) {
		t.Errorf("expected %v to contain %v", input, *selected)
	}
}

func TestSliceSelectRandomWithEmptySlice(t *testing.T) {
	selected := SliceSelectRandom([]string{})

	if selected != nil {
		t.Errorf("expected nil, got %v", selected)
	}
}

func TestSliceSelectRandomWithGeneratorWithNilGenerator(t *testing.T) {
	selected := SliceSelectRandomWithGenerator([]string{"a", "b"}, nil)

	if selected != nil {
		t.Errorf("expected nil, got %v", selected)
	}
}
