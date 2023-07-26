package brain

import (
	"errors"
)

// NetInput calculates the linear combination of two vectors with bias.
func NetInput(inputs []float64, weights [][]float64, bias []float64) ([]float64, error) {

	if len(weights) != len(bias) {
		return nil, errors.New("Error: Size of weights array does not match size of bias.")
	}

	outputs := []float64{}

	for i := range weights {
		if len(inputs) != len(weights[i]) {
			return nil, errors.New("Error: Size of input vector does not match size of weights.")
		}

		output := bias[i]

		for j := range inputs {
			output += inputs[j] * weights[i][j]
		}

		outputs = append(outputs, output)

	}

	return outputs, nil
}
