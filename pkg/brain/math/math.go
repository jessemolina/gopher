package math

import (
	"math/rand"
)

// RandomSlice64 generates a random float64 slice.
func RandomSlice64(size int) []float64 {
	slice := make([]float64, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Float64()
	}
	return slice
}
