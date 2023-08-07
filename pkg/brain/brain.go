package brain

import (
	"errors"

	"github.com/jessemolina/gopher/pkg/brain/math"
)

/*
   TODO Standardize return error messages and incorporate them in table tests.
   TODO Consider creating global options type for passing a seed to random builder functions.
*/

// Neuron is a single perceptron node.
type Neuron struct {
	weights []float64
	bias    float64
}

// WeightedSum calculates the linear combination of inputs and neuron weights with bias.
func (n *Neuron) WeightedSum(inputs []float64) (float64, error) {
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

// ForwardPass calculates the WeightedSum of all neurons in a layer.
func (l *Layer) ForwardPass(inputs [][]float64) ([][]float64, error) {
	output := [][]float64{}
	for _, input := range inputs {
		results := []float64{}
		for _, neuron := range l.Neurons {
			netOutput, err := neuron.WeightedSum(input)
			if err != nil {
				return nil, err
			}
			results = append(results, netOutput)
		}

		output = append(output, results)
	}
	return output, nil
}

// DenseLayer creates a new dense layer with random neuron weights and biases.
func DenseLayer(inputs, neurons int) (*Layer, error) {
	if inputs == 0 || neurons == 0 {
		return nil, errors.New("Error: Invalid number of inputs or neurons.")
	}

	ns := []Neuron{}
	for i := 0; i < neurons; i++ {
		n := Neuron{math.RandomSlice64(inputs), 0}
		ns = append(ns, n)
	}

	return &Layer{ns}, nil
}
