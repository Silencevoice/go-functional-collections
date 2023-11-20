package collections

import "testing"

func TestSortInteger(t *testing.T) {
	slice := []int{4, 2, 0, 1, 3}
	expect := []int{0, 1, 2, 3, 4}
	got := Sort(slice)

	for i := 0; i < len(expect); i++ {
		if expect[i] != got[i] {
			t.Errorf("Position %d of sorted slice is not %v", i, expect[i])
		}
	}
}

func TestSortFloat(t *testing.T) {
	slice := []float64{0.4, 0.2, 0.0, 0.1, 0.3}
	expect := []float64{0.0, 0.1, 0.2, 0.3, 0.4}
	got := Sort(slice)

	for i := 0; i < len(expect); i++ {
		if expect[i] != got[i] {
			t.Errorf("Position %d of sorted slice is not %v", i, expect[i])
		}
	}
}
