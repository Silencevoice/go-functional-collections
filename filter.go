package collections

// KeepFunc is an alias to name any function that takes a E parameter and returns a boolean.
type KeepFunc[E any] func(E) bool

// Filter applies a KeepFunc to every elemento of the slice an creates a new slice with only the elementos for which f returned true.
func Filter[E any](slice []E, f KeepFunc[E]) []E {
	result := []E{}
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
