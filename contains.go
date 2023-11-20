package collections

// Contains checks whether a value is present in a slice.
func Contains[E comparable](slice []E, value E) bool {
	for _, e := range slice {
		if e == value {
			return true
		}
	}

	return false
}
