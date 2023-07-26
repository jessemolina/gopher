package brain

import (
	"fmt"
	"testing"
)

// message is a formatted error for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// tests is a table used for table testing different functions.
var tests = []struct {
	inputs   []float64
	weights  [][]float64
	bias     []float64
	expected []float64
}{
	{
		[]float64{1, 2, 3},
		[][]float64{{0.2, 0.8, -0.5}},
		[]float64{2.0},
		[]float64{2.3000000000000003}},
	{
		[]float64{1.0, 2.0, 3.0, 2.5},
		[][]float64{{0.2, 0.8, -0.5, 1.0}},
		[]float64{2.0},
		[]float64{4.800000000000001}},
	{
		[]float64{1.0, 2.0, 3.0, 2.5},
		[][]float64{{0.2, 0.8, -0.5, 1.0}, {0.5, -0.91, 0.26, -0.5}, {-0.26, -0.27, 0.17, 0.87}},
		[]float64{2.0, 3.0, 0.5},
		[]float64{4.800000000000001, 1.21, 2.385}},
}

// TestNetInput tests the NetInput function.
func TestNetInput(t *testing.T) {
	for _, test := range tests {
		results, err := NetInput(test.inputs, test.weights, test.bias)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(results)

		for i := range results {
			if results[i] != test.expected[i] {
				t.Errorf(message, test.expected[i], results[i])
			}
		}

	}
}
