package brain

import (
	"math/rand"
	"testing"

	"github.com/jessemolina/gopher/pkg/brain/data"
)

const (
	seed = 1000
)

// message is a formatted error for failed tests.
var message = "\nTest:\t\t %v\nExpected:\t %v\nResults:\t %v\n"

// TODO Consider creating the layers dynamically and testing based on specific seed results.
// tests is a collection of table tests with the expected result.
var tests = []struct {
	layer    Layer
	inputs   [][]float64
	expected []float64
}{
	{
		layer: Layer{
			[]Neuron{
				{weights: []float64{0.2, 0.8, -0.5}, bias: 2.0},
			},
		},
		inputs:   [][]float64{{1, 2, 3}},
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
		inputs:   [][]float64{{1, 2, 3, 2.5}},
		expected: []float64{4.800000000000001, 1.21, 2.385},
	},
	{
		layer: Layer{
			[]Neuron{
				{weights: []float64{0.6806635511317619, 0.28981398417996873, 0.5357723685986947}, bias: 0},
			},
		},
		inputs:   [][]float64{{1, 2, 3}},
		expected: []float64{2.8676086252877835},
	},
	/*{
		layer: *DenseLayer(3, 1),
		inputs:   []float64{1, 2, 3},
		expected: []float64{2.8676086252877835},
	},*/
}

// init is used to initialize pre-requisites for the tests.
func init() {
	rand.Seed(seed)
}

// Test the Neuron.NetInput method.
func TestNetInput(t *testing.T) {
	for i, test := range tests {
		results, err := test.layer.Neurons[0].NetInput(test.inputs[0])
		if err != nil {
			t.Fatal(err)
		}

		/*
		fmt.Println("\n", i)
		fmt.Println("\t", test.layer)
		fmt.Println("\t", results)
		*/

		if results != test.expected[0] {
			t.Errorf(message, i, test.expected[0], results)

		}
	}
}

// Test the Layer.WeightedSum method.
func TestWeightedSum(t *testing.T) {
	for i, test := range tests {
		results, err := test.layer.WeightedSum(test.inputs)
		if err != nil {
			t.Fatal(err)
		}

		for j := range results {
			if results[j][0] != test.expected[j] {
				t.Errorf(message, i, test.expected, results)
			}
		}
	}
}

// TODO Create table test for TestDenseLayer.
// Test the DenseLayer function.
func TestDenseLayer(t *testing.T) {
	inputs, neurons := 0, 1
	results, err := DenseLayer(inputs, neurons)
	if err != nil {
		errInvalidNumber := "Error: Invalid number of inputs or neurons."
		if err.Error() != errInvalidNumber  {
			t.Errorf(message, 0, err.Error(), errInvalidNumber)
		}
	}

	if results == nil {
		t.SkipNow()
	}

	if neurons != len(results.Neurons) {
		t.Errorf(message, 0, neurons, len(results.Neurons))
	}

	if inputs != len(results.Neurons[0].weights) {
		t.Errorf(message, 0, inputs, len(results.Neurons[0].weights))
	}
}

// TestDenseLayerForward tests the weighted sum for a generated data set.
func TestDenseLayerForward(t *testing.T) {
	samples := 100
	cardinality := 3

	dataset := data.SpiralDataset(samples, cardinality)
	layer, err := DenseLayer(2, 3)
	if err != nil {
		t.Fatal(err)
	}

	results, err := layer.WeightedSum(dataset.X)
	if err != nil {
		t.Fatal(err)
	}

	if samples != len(results) {
		t.Errorf(message, 0, samples, len(results))
	}
}
