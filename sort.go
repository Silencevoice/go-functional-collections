package collections

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort sorts a slice, provided its type is Ordered.
func Sort[E constraints.Ordered](slice []E) []E {
	result := make([]E, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
