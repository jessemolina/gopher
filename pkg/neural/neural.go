package neural

import (
	"errors"
)

// NetInput calculates the linear combination of two vectors with bias.
func NetInput(inputs, weights []float64, bias float64) (float64, error) {
	if len(inputs) != len(weights){
		return 0, errors.New("Error: Vector dimmension mismatch.")
	}

	output := bias
	for i := range inputs{
		output += inputs[i] * weights[i]
	}

   return output, nil
}
