package collections

import (
	"strconv"
	"testing"
)

func TestReverseStringSlice(t *testing.T) {
	slice := []string{"Hello", "World", "This", "Is", "A", "Test"}

	expect := []string{"Test", "A", "Is", "This", "World", "Hello"}
	got := Reverse(slice)

	if len(expect) != len(got) {
		t.Errorf("Reverse of %v has diffrent length. got %d, expected %d.", slice, len(got), len(expect))
	}

	// check evey item in expect is in got
	for i := 0; i < len(expect); i++ {
		if expect[i] != got[i] {
			t.Errorf("Reverse position %d of slice %v is not %s", i, slice, expect[i])
		}
	}
}

type testStruct struct {
	intValue   int
	intPts     *int
	strValue   string
	strPointer *string
}

func newTestStruct(intValue int, strValue string) testStruct {
	return testStruct{
		intValue:   intValue,
		intPts:     &intValue,
		strValue:   strValue,
		strPointer: &strValue,
	}
}

func TestReverseStructSlice(t *testing.T) {
	slice := []testStruct{}
	for i := 0; i < 10; i++ {
		intValue := i
		slice = append(slice, newTestStruct(intValue, "str "+strconv.Itoa(intValue)))
	}
	reverse := Reverse(slice)

	for i := 0; i < 10; i++ {
		expect := slice[9-i]
		got := reverse[i]
		if expect != got {
			t.Errorf("Reverse position %d is not %v", i, expect)
		}
	}
}
