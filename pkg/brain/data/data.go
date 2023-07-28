package data

import (
	"math"
	"math/rand"
)

// Sample represents a single record in a dataset with predictors and target values.
type Sample struct {
	Predictors []float64
	Target     float64
}

// SpiralDataset creates a dataset of X, Y spiral coordinates with a categorical label.
// (x, y) = (cos t, sin t), where t is the angle in radians.
func SpiralDataset(samples, cardinality int) []Sample {
	count := samples / cardinality
	remaining := int(samples % cardinality)

	data := []Sample{}

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

		for j := 0; j < count ; j++ {
			scale := float64(j) / float64(count)
			sample := Sample{xy(scale, ordinal), ordinal}

			data = append(data, sample)
		}
	}

	// Add remainder count of sample data points to a random cardinal class.
	for i := 0; i < remaining ; i++ {
		ordinal := float64(rand.Intn(4))
		scale := float64(i)/ float64(remaining)
		sample := Sample{xy(scale, ordinal), float64(i)}

		data = append(data, sample)
	}

	return data
}
