package data

import (
	"testing"
)

// message is a formatted error for failed tests.
var message = "\nTest:\t\t %v\nExpected:\t %v\nResults:\t %v\n"

// TestSpiralDataset validates that proper sample and class count are generated in the dataset.
func TestSpiralDataset(t *testing.T) {
	samples := 100
	cardinality := 3

	data := SpiralDataset(samples, cardinality)

	unique := make(map[float64]bool)

	for _, sample := range data {
		unique[sample.Target] = true
	}

	if cardinality != len(unique) {
		t.Errorf(message, 0, cardinality, len(unique))
	}

	if samples != len(data) {
		t.Errorf(message, 0, samples, len(data))
	}
}
