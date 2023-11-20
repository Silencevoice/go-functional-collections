package collections

// Alias to name any function that takes a E parameter and returns a parameter of type T.
type MapFunc[E any, T any] func(E) T

// Map maps every element of the slice applying the function f to every elemento.
func Map[E any, T any](slice []E, f MapFunc[E, T]) []T {
	result := make([]T, len(slice))
	for i := range slice {
		result[i] = f(slice[i])
	}
	return result
}
