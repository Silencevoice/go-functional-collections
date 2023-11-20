package collections

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	str := "Hello world"
	slice := []string{str, ""}

	expect := true
	got := Contains(slice, str)
	if expect != got {
		t.Errorf("Failed. slice %v should contain %s string", slice, str)
	}

	expect = false
	got = Contains(slice, "Not contains")
	if expect != got {
		t.Errorf("Failed. slice %v should not contain %s string", slice, str)
	}
}

func TestContainsInt(t *testing.T) {
	value := 5
	slice := []int{value, 2, 0, -1}

	expect := true
	got := Contains(slice, value)
	if expect != got {
		t.Errorf("Failed. slice %v should contain %d int", slice, value)
	}

	expect = false
	got = Contains(slice, -5)
	if expect != got {
		t.Errorf("Failed. slice %v should not contain %d int", slice, value)
	}
}

func TestContainsNilSlice(t *testing.T) {
	value := "Hello"
	var slice []string = nil

	expect := false
	got := Contains(slice, value)

	if got != expect {
		t.Errorf("Failed. slice %v should not contain %s string", slice, value)
	}
}
