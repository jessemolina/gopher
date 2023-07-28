package data

import (
	"math"
	"math/rand"
)

// Dataset represents a collection of X inputs and Y targets.
type Dataset struct {
	X [][]float64
	Y []float64
}

// SpiralDataset creates a dataset of X, Y spiral coordinates with a categorical label.
// (x, y) = (cos t, sin t), where t is the angle in radians.
func SpiralDataset(samples, cardinality int) Dataset {
	count := samples / cardinality
	remaining := int(samples % cardinality)

	dataset := Dataset{}

	// Anonymous function to generate random coordinates of a circle.
	xy := func(scale, ordinal float64) []float64 {
		angle := (math.Pi * 2.0) * scale
		radius := (rand.Float64() * 0.5) + (ordinal + 0.5)

		x := radius * math.Cos(angle)
		y := radius * math.Sin(angle)

		return []float64{x, y}
	}

	// Add evenly-split count of sample data points per cardinal class.
	for i := 0; i < cardinality; i++ {
		ordinal := float64(i)

		for j := 0; j < count; j++ {
			scale := float64(j) / float64(count)
			sample := xy(scale, ordinal)

			dataset.X = append(dataset.X, sample)
			dataset.Y = append(dataset.Y, ordinal)
		}
	}

	// Add remainder count of sample data points to a random cardinal class.
	for i := 0; i < remaining; i++ {
		ordinal := float64(rand.Intn(cardinality))
		scale := float64(i) / float64(remaining)
		sample := xy(scale, ordinal)

		dataset.X = append(dataset.X, sample)
		dataset.Y = append(dataset.Y, ordinal)
	}

	return dataset
}
