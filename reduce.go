package collections

// ReduceFunc is an alias for any function that takes two elements of the same type E and performs an operation that transforms then into a result of the type T.
type ReduceFunc[E any, T any] func(T, E) T

// Reduce applies the ReduceFunc[E] to the elements of a slice of type E, performing the operation in sequence and thus reducing the slice to a single value.
func Reduce[E any, T any](slice []E, init T, f ReduceFunc[E, T]) T {
	cur := init
	for _, v := range slice {
		cur = f(cur, v)
	}
	return cur
}
