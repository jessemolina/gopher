package data

import (
	"testing"
)

// message is a formatted error for failed tests.
var message = "\nTest:\t\t %v\nExpected:\t %v\nResults:\t %v\n"

// TODO Create table test for TestSpiralDataset
// TestSpiralDataset validates that proper sample and class count are generated in the dataset.
func TestSpiralDataset(t *testing.T) {
	samples := 100
	cardinality := 3

	data := SpiralDataset(samples, cardinality)

	unique := make(map[float64]bool)

	for _, sample := range data.Y {
		unique[sample] = true
	}

	if cardinality != len(unique) {
		t.Errorf(message, 0, cardinality, len(unique))
	}

	if samples != len(data.X) {
		t.Errorf(message, 0, samples, len(data.X))
	}
}
