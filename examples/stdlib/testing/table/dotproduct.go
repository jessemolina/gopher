package main

// Array of scalars.
type vector []float64

// Multiplies two vectors and returns a scalar value.
// a1*b1 + a2*b2 + ... + an*bn
func DotProduct(a vector, b vector) float64 {
	s := 0.0
	if len(a) == len(b) {
		for i := range a {
			s += a[i] * b[i]
		}
	}
	return s
}
