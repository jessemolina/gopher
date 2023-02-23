package main

import "errors"

// Array of scalars.
type vector []float64

// Multiplies two vectors and returns a scalar value.
// a1*b1 + a2*b2 + ... + an*bn
func DotProduct(a vector, b vector) (float64, error) {
	s := 0.0

	// Verify that vectors contain the same number of scalars.
	if len(a) != len(b) {
		return s, errors.New("Vector dimmensions don't match")
	}

	// Multiply corresponding components and add.
	for i := range a {
		s += a[i] * b[i]
	}

	return s, nil
}
