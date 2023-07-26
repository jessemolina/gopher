package neural

import (
	"testing"
)

// message is a formatted error for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// tests is a table used for table testing different functions.
var tests = []struct{
	inputs []float64
	weights []float64
	bias float64
	expected float64
}{
	{[]float64{1, 2, 3}, []float64{0.2, 0.8, -0.5}, 2.0, 2.3000000000000003},
	{[]float64{1.0, 2.0, 3.0, 2.5}, []float64{0.2,0.8,-0.5, 1.0}, 2.0, 4.800000000000001},
}

// TestNetInput tests the NetInput function.
func TestNetInput(t *testing.T) {
	for _ , test := range tests {
		results, _ := NetInput(test.inputs, test.weights, test.bias)

		if results != test.expected {
			t.Errorf(message, test.expected, results)
		}
	}
}
