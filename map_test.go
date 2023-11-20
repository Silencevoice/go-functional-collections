package collections

import (
	"strings"
	"testing"
)

func TestMapStringToInt(t *testing.T) {
	slice := []string{"Hello", "World"}
	var f MapFunc[string, int] = func(str string) int {
		return len(str)
	}

	expected := []int{5, 5}
	got := Map(slice, f)

	for i, exp := range expected {
		g := got[i]
		if g != exp {
			t.Errorf("Element %d of slice %v was not mapped as expected. Got %d and expected %d", i, slice, g, exp)
		}
	}
}

var toUpper MapFunc[string, string] = func(str string) string {
	return strings.ToUpper(str)
}

func TestMapStringToString(t *testing.T) {
	slice := []string{"Hello", "World"}

	expected := []string{"HELLO", "WORLD"}
	got := Map(slice, toUpper)

	for i, exp := range expected {
		g := got[i]
		if g != exp {
			t.Errorf("Element %d of slice %v was not mapped as expected. Got %s and expected %s", i, slice, g, exp)
		}
	}
}

// Adding method to slice type
type myType []string

func (s myType) MapToString(f MapFunc[string, string]) []string {
	return Map(s, f)
}

func TestMyTypeMapToString(t *testing.T) {
	slice := []string{"Hello", "World"}
	typeSlice := myType(slice)

	expected := []string{"HELLO", "WORLD"}
	got := typeSlice.MapToString(toUpper)

	for i, exp := range expected {
		g := got[i]
		if g != exp {
			t.Errorf("Element %d of slice %v was not mapped as expected. Got %s and expected %s", i, slice, g, exp)
		}
	}
}
