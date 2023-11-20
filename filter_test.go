package collections

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// IsEven is a generic function that returns true if value is event, false if it's not.
func IsEven[T constraints.Integer](value T) bool {
	return value%2 == 0
}

func TestIsEvenInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	expected := []int{2, 4, 6}
	got := Filter(slice, IsEven)

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("unexpected elemenmt after filtering at position %d. Expected %d. Got %d.", i, e, got[i])
		}
	}
}

type IntSlice []int

func (is IntSlice) Filter(f KeepFunc[int]) IntSlice {
	return Filter(is, f)
}

func TestIntSliceIsEven(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	intSlice := IntSlice(slice)

	expected := []int{2, 4, 6}
	got := intSlice.Filter(IsEven)

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("unexpected elemenmt after filtering at position %d. Expected %d. Got %d.", i, e, got[i])
		}
	}
}

func TestIntSliceIsEvenInline(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	intSlice := IntSlice(slice)

	expected := []int{2, 4, 6}
	got := intSlice.Filter(func(i int) bool {
		return i%2 == 0
	})

	for i, e := range expected {
		if e != got[i] {
			t.Errorf("unexpected elemenmt after filtering at position %d. Expected %d. Got %d.", i, e, got[i])
		}
	}
}
