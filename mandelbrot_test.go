package mandelbrot

import "testing"

func TestConfigImaginaryMax(t *testing.T) {
	tests := []struct {
		width    int
		height   int
		expected float64
	}{
		{1000, 1000, 2.0},
		{1000, 750, 1.0},
		{1000, 500, 0.0},
	}
	for _, test := range tests {
		config := NewConfig(test.width, test.height, 20)
		max := config.imaginaryMax
		if max != test.expected {
			t.Errorf("Incorrect imaginaryMax, got: %f, expected: %f", max, test.expected)
		}
	}
}

func TestCoordRealAndImaginary(t *testing.T) {
	tests := []struct {
		x        int
		y        int
		expected complex128
	}{
		{0, 0, complex(-2.0, 2.0)},
		{999, 999, complex(2.0, -2.0)},
	}
	config := NewConfig(1000, 1000, 10)
	for _, test := range tests {
		coord := Coord{test.x, test.y}
		result := coord.complex(config)
		if result != test.expected {
			t.Errorf("Incorrect complex value, got: %f, expected: %f", result, test.expected)
		}
	}
}
