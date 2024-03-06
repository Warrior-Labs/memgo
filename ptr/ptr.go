package ptr

// To returns a pointer to the given value
func To[T any](v T) *T {
	return &v
}

// From returns the value pointed to by the given pointer
func From[T any](p *T) T {
	return *p
}
