package brain

import (
	"errors"
)

// Neuron is a single perceptron node.
type Neuron struct {
	weights []float64
	bias    float64
}

// NetInput calculates the linear combination with bias of a nueron.
func (n *Neuron) NetInput(inputs []float64) (float64, error) {
	if len(inputs) != len(n.weights) {
		return 0, errors.New("Error: Mismatch in size of neuron inputs and weights.")
	}

	output := n.bias
	for i := range inputs {
		output += inputs[i] * n.weights[i]
	}

	return output, nil
}

// Layer is a collection of neurons.
type Layer struct {
	Neurons []Neuron
}

// WeightedSum calculates the linear cobination of inputs and neurons in a layer.
func (l *Layer) WeightedSum(inputs []float64) ([]float64, error) {
	output := []float64{}
	for _, n := range l.Neurons {
		results, err := n.NetInput(inputs)
		if err != nil {
			return nil, err
		}
		output = append(output, results)
	}

	return output, nil
}
