package collections

import "testing"

func TestReduceIntToIntSum(t *testing.T) {
	slice := []int{1, 2, 3, 4}

	expected := 10
	got := Reduce(slice, 0, func(cur, next int) int {
		return cur + next
	})

	if expected != got {
		t.Errorf("Reduce value error. Expected %d. Got %d.", expected, got)
	}
}

var fnSum = func(cur, next int) int {
	return cur + next
}

func TestReduceIntToIntExternalSum(t *testing.T) {
	slice := []int{1, 2, 3, 4}

	expected := 10
	got := Reduce(slice, 0, fnSum)

	if expected != got {
		t.Errorf("Reduce value error. Expected %d. Got %d.", expected, got)
	}
}

func (is IntSlice) ReduceToInt(fn ReduceFunc[int, int]) int {
	return Reduce(is, 0, fn)
}

func TestReduceIntSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	intSlice := IntSlice(slice)
	expected := 10
	got := intSlice.ReduceToInt(fnSum)

	if expected != got {
		t.Errorf("Reduce value error. Expected %d. Got %d.", expected, got)
	}
}
