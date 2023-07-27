package brain

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	seed = 1000
)

// message is a formatted error for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// TODO Consider creating the layers dynamically and testing based on specific seed results.
// tests is a collection of table tests with the expected result.
var tests = []struct {
	layer    Layer
	inputs   []float64
	expected []float64
}{
	{
		layer: Layer{
			[]Neuron{
				{weights: []float64{0.2, 0.8, -0.5}, bias: 2.0},
			},
		},
		inputs:   []float64{1, 2, 3},
		expected: []float64{2.3000000000000003},
	},
	{
		layer: Layer{
			[]Neuron{
				{weights: []float64{0.2, 0.8, -0.5, 1}, bias: 2.0},
				{weights: []float64{0.5, -0.91, 0.26, -0.5}, bias: 3.0},
				{weights: []float64{-0.26, -0.27, 0.17, 0.87}, bias: 0.5},
			},
		},
		inputs:   []float64{1, 2, 3, 2.5},
		expected: []float64{4.800000000000001, 1.21, 2.385},
	},
}

// init is used to initialize pre-requisites for the tests.
func init() {
	rand.Seed(seed)
}

// Test the Neuron.NetInput method.
func TestNetInput(t *testing.T) {
	for _, test := range tests {
		results, err := test.layer.Neurons[0].NetInput(test.inputs)
		if err != nil {
			t.Fatal(err)
		}

		if results != test.expected[0] {
			t.Errorf(message, test.expected[0], results)

		}
	}
}

// Test the Layer.WeightedSum method.
func TestWeightedSum(t *testing.T) {
	for _, test := range tests {
		results, err := test.layer.WeightedSum(test.inputs)
		if err != nil {
			t.Fatal(err)
		}

		for i := range results {
			if results[i] != test.expected[i] {
				t.Errorf(message, test.expected, results)
			}
		}
	}
}

// TODO Create table test for TestDenseLayer.
// Test the DenseLayer function.
func TestDenseLayer(t *testing.T) {
	inputs, neurons := 3, 5
	results := DenseLayer(inputs, neurons)

	fmt.Println(results)

	if neurons != len(results.Neurons) {
		t.Errorf(message, neurons, len(results.Neurons))
	}

	if inputs != len(results.Neurons[0].weights) {
		t.Errorf(message, inputs, len(results.Neurons[0].weights))
	}
}
