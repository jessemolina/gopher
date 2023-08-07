package activate

// Linear is a Linear activation functions
func Linear(input float64) float64 {
	return input
}

// ReLU is a Rectified Linear Unit activation function.
func ReLU(input float64) float64 {
	if input > 0 {
		return input
	}

	return 0
}

