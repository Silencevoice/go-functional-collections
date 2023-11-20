package collections

// Reverse reverses any slice.
func Reverse[E any](slice []E) []E {
	result := make([]E, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}
